import grpc
import uuid
import json
import requests
from fastavro import parse_schema
import beneath
from beneath.stream import Stream
from beneath.proto import engine_pb2
from beneath.proto import gateway_pb2
from beneath.proto import gateway_pb2_grpc
from beneath import config
  
# create a map on the client. instanceid -> schema, so we cache it to remember it
# create fn "getAvroSchema" to see if the schema is in memory already

class Client:
  """Client to bundle configuration for API requests.

  Args:
    secret (str):
      The user's password to authenticate permission to access Beneath. 
  """

  # initialize the client with the user's secret
  def __init__(self, secret):
    self.secret = secret
    if not isinstance(secret, str):
      raise TypeError("secret must be a string")
    
    self._prepare()

  def __getstate__(self):
    return {
      "secret": self.secret,
    }  

  def __setstate__(self, obj):
    self.secret = obj["secret"]
    self._prepare()

  # create a client with the provided secret
  def _prepare(self):
    self.request_metadata = [('authorization', 'Bearer {}'.format(self.secret))]

    # open a grpc channel from the client to the server
    # TODO: create a SSL/TLS connection
    self.channel = grpc.insecure_channel('localhost:50051')

    # create a "stub" (aka a client). the stub has all the methods that the gateway server has. so it'll have ReadRecords(), WriteRecords(), and GetStreamDetails()
    self.stub = gateway_pb2_grpc.GatewayStub(self.channel)

    # ensure that the user is running the most current Python package
    response = self.stub.GetCurrentBeneathPackageVersion(
        gateway_pb2.PackageVersionRequest(package_version=beneath.__version__),
      metadata=self.request_metadata)
    if response.version_response == "not current":
      raise Exception(
          "Your Beneath package is not up-to-date. Please upgrade before continuing.")

    # create a dictionary to remember schemas
    self.avro_schemas = dict()

  # get a stream's details
  def stream(self, project_name, stream_name):
    details = self.stub.GetStreamDetails(
        gateway_pb2.StreamDetailsRequest(
            project_name=project_name, stream_name=stream_name),
        metadata=self.request_metadata)

    # store the stream's schema in memory
    self.avro_schemas[details.current_instance_id] = details.avro_schema

    # return a Stream class
    return Stream(
      client=self,
      project_name=details.project_name,
      stream_name=details.stream_name,
      current_instance_id=uuid.UUID(bytes=details.current_instance_id),
      avro_schema=parse_schema(json.loads(details.avro_schema)),
      batch=details.batch,
    )

  # Client code for control server
  # run a GraphQL query
  def run_query(self, query, variables):
    headers = {"Authorization": "Bearer " + self.secret}
    request = requests.post(config.BENEATH_CONTROL_HOST + '/graphql', json={'query': query, 'variables': variables}, headers=headers)
    if request.status_code == 200:
      return request.json()
    else:
      print(request.text)
      raise Exception("Query failed to run by returning code of {}. {}".format(request.status_code, query))

  # get "me" info
  def get_me(self):
    result = self.run_query(
      variables={},
      query="""
      query Me {
          me{
            userID
            user {
              username
              name
            }
            email
            updatedOn
          }
        }
      """
    )
    return result['data']['me']

  # get a user's info by UUID
  def get_user_by_ID(self, userID):
    result = self.run_query(
      variables={
          "userID": userID
      },
      query="""
      query User($userID: UUID!) {
        user(
          userID: $userID
        ) {
          userID
          username
          name
          bio
          photoURL
          createdOn
          projects {
            name
            createdOn
            updatedOn
            streams {
              name
            }
          }
        }
      }
      """
    )
    return result['data']['user']

  # get a project's info by name
  def get_project_by_name(self, project_name):
    result = self.run_query(
        variables={
            "name": project_name
        },
        query="""
        query ProjectByName($name: String!) {
            projectByName(name: $name) {
                projectID
                name
                displayName
                site
                description
                photoURL
                createdOn
                updatedOn
                users {
                  username
                }
                streams {
                  name
                }
            }
        }
      """
    )
    return result['data']['projectByName']

  # get a streams's info by stream name and project name
  def get_stream_details(self, stream_name, project_name):
    result = self.run_query(
        variables={
            "name": stream_name,
            "projectName": project_name
        },
        query="""
        query Stream($name: String!, $projectName: String!) {
            stream(
              name: $name, 
              projectName: $projectName) {
                streamID
                name
                description
                schema
                avroSchema
                keyFields
                external
                batch
                manual
                project {
                  name
                }
                currentStreamInstanceID
                createdOn
                updatedOn
            }
        }
      """
    )
    return result['data']['stream']

  # create an external stream
  def create_external_stream(self, project_id, schema, manual):
    result = self.run_query(
        variables={
            "projectID": project_id,
            "schema": schema,
            "batch": False,
            "manual": manual if manual else False
        },
        query="""
        mutation CreateExternalStream($projectID: UUID!, $schema: String!, $batch: Boolean!, $manual: Boolean!) {
          createExternalStream(
            projectID: $projectID,
            schema: $schema,
            batch: $batch,
            manual: $manual
          ) {
            streamID
            name
            description
            schema
            avroSchema
            keyFields
            external
            batch
            manual
            project {
              projectID
              name
            }
            currentStreamInstanceID
            createdOn
            updatedOn
          }
        }
      """
    )
    return result

  # update an external stream
  def update_external_stream(self, stream_id, schema, manual):
    result = self.run_query(
        variables={
            "streamID": stream_id,
            "schema": schema,
            "manual": manual if manual else False
        },
        query="""
        mutation UpdateStream($streamID: UUID!, $schema: String!, $manual: Boolean!) {
          updateStream(
            streamID: $streamID,
            schema: $schema,
            manual: $manual
          ) {
            streamID
            name
            description
            schema
            avroSchema
            keyFields
            external
            batch
            manual
            project {
              projectID
              name
            }
            currentStreamInstanceID
            createdOn
            updatedOn
          }
        }
      """
    )
    return result
