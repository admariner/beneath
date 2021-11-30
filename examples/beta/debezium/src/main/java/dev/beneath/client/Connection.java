package dev.beneath.client;

import java.io.IOException;
import java.util.UUID;

import com.apollographql.apollo.ApolloClient;
import com.google.protobuf.ByteString;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import dev.beneath.client.utils.Utils;
import io.grpc.CallOptions;
import io.grpc.Channel;
import io.grpc.ClientCall;
import io.grpc.ClientInterceptor;
import io.grpc.ForwardingClientCall;
import io.grpc.ManagedChannelBuilder;
import io.grpc.Metadata;
import io.grpc.Metadata.Key;
import io.grpc.MethodDescriptor;
import okhttp3.Interceptor;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;

/**
 * Encapsulates network connectivity to Beneath
 */
public class Connection {
  private String secret;
  private Boolean connected;
  private Metadata requestMetadata;
  private Channel channel;
  private GatewayGrpc.GatewayBlockingStub blockingStub;
  // TODO: make relevant functions async
  private GatewayGrpc.GatewayStub asyncStub;
  private PingResponse pong;
  public ApolloClient apolloClient;

  private static final Logger LOGGER = LoggerFactory.getLogger(Connection.class);

  public Connection(String secret) {
    this.secret = secret;
    this.connected = false;

    // construct metadata
    this.requestMetadata = new Metadata();
    this.requestMetadata.put(Metadata.Key.of("authorization", Metadata.ASCII_STRING_MARSHALLER),
        String.format("Bearer %s", this.secret));
  }

  // GRPC CONNECTIVITY

  public class DataPlaneAuthInterceptor implements ClientInterceptor {
    public <ReqT, RespT> ClientCall<ReqT, RespT> interceptCall(final MethodDescriptor<ReqT, RespT> methodDescriptor,
        final CallOptions callOptions, final Channel channel) {
      return new ForwardingClientCall.SimpleForwardingClientCall<ReqT, RespT>(
          channel.newCall(methodDescriptor, callOptions)) {
        @Override
        public void start(final Listener<RespT> responseListener, final Metadata headers) {
          headers.put(Key.of("Authorization", Metadata.ASCII_STRING_MARSHALLER), String.format("Bearer %s", secret));
          super.start(responseListener, headers);
        }
      };
    }
  }

  public void ensureConnected() {
    if (!connected) {
      createGrpcConnection("host.docker.internal", 50051);
      PingResponse pong = ping();
      checkPongStatus(pong);
      this.pong = pong;
      connected = true;
    }
    if (!this.pong.getAuthenticated()) {
      throw new AuthenticationException("You must authenticate with 'beneath auth' or by setting BENEATH_SECRET");
    }
  }

  private void createGrpcConnection(String host, Integer port) {
    Boolean insecure = true;
    if (insecure) {
      this.channel = ManagedChannelBuilder.forAddress(host, port).usePlaintext()
          .intercept(new DataPlaneAuthInterceptor()).build();
    } else {
      this.channel = ManagedChannelBuilder.forAddress(host, port).intercept(new DataPlaneAuthInterceptor()).build();
    }
    this.blockingStub = GatewayGrpc.newBlockingStub(this.channel).withCompression("gzip");
    this.asyncStub = GatewayGrpc.newStub(this.channel).withCompression("gzip");
    this.connected = true;
  }

  private static void checkPongStatus(PingResponse pong) {
    if (Config.DEV) {
      return;
    }
    if (pong.getVersionStatus() == "warning") {
      LOGGER.warn("This version ({}) of the Beneath java library will soon be deprecated (recommended: {}).",
          Config.JAVA_CLIENT_VERSION, pong.getRecommendedVersion());
    } else if (pong.getVersionStatus() == "deprecated") {
      throw new RuntimeException(
          String.format("This version (%s) of the Beneath java library is out-of-date (recommended: %s).",
              Config.JAVA_CLIENT_VERSION, pong.getRecommendedVersion()));
    }
  }

  private PingResponse ping() {
    PingRequest request = PingRequest.newBuilder().setClientId(Config.JAVA_CLIENT_ID)
        .setClientVersion(Config.JAVA_CLIENT_VERSION).build();

    PingResponse response = this.blockingStub.ping(request);

    return response;
  }

  // CONTROL PLANE

  private class ControlPlaneAuthInterceptor implements Interceptor {
    @Override
    public Response intercept(Interceptor.Chain chain) throws IOException {
      Request request = chain.request().newBuilder().addHeader("Authorization", String.format("Bearer %s", secret))
          .build();
      return chain.proceed(request);
    }
  }

  public void createGraphQlConnection() {
    this.apolloClient = ApolloClient.builder().serverUrl(Config.BENEATH_CONTROL_HOST)
        .okHttpClient(new OkHttpClient.Builder().addInterceptor(new ControlPlaneAuthInterceptor()).build()).build();
  }

  // DATA PLANE

  public WriteResponse write(InstanceRecords instanceRecords) {
    this.ensureConnected();
    WriteRequest request = WriteRequest.newBuilder().addInstanceRecords(instanceRecords).build();
    return this.blockingStub.write(request);
  }

  public QueryLogResponse queryLog(UUID instanceId, Boolean peek) {
    this.ensureConnected();
    ByteString instanceIdByteString = ByteString.copyFrom(Utils.uuidToBytes(instanceId));
    QueryLogRequest request = QueryLogRequest.newBuilder().setInstanceId(instanceIdByteString).setPartitions(1)
        .setPeek(peek).build();
    return this.blockingStub.queryLog(request);
  }

  public QueryIndexResponse queryIndex(UUID instanceId, String filter) {
    this.ensureConnected();
    ByteString instanceIdByteString = ByteString.copyFrom(Utils.uuidToBytes(instanceId));
    QueryIndexRequest request = QueryIndexRequest.newBuilder().setInstanceId(instanceIdByteString).setPartitions(1)
        .setFilter(filter).build();
    return this.blockingStub.queryIndex(request);
  }

  public ReadResponse read(ByteString cursor, Integer limit) {
    this.ensureConnected();
    ReadRequest request = ReadRequest.newBuilder().setCursor(cursor).setLimit(limit).build();
    return this.blockingStub.read(request);
  }
}