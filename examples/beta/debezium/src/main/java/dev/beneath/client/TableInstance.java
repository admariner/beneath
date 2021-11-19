package dev.beneath.client;

import java.util.UUID;

import com.google.protobuf.ByteString;

import dev.beneath.CreateTableMutation;
import dev.beneath.TableByOrganizationProjectAndNameQuery;
import dev.beneath.CreateTableInstanceMutation.CreateTableInstance;

/**
 * Represents an instance of a Table, i.e. a specific version that you can
 * query/subscribe/write to. Learn more about instances at
 * https://about.beneath.dev/docs/concepts/tables/.
 */
public class TableInstance {
  public Table table;
  public UUID instanceId;
  public Boolean isFinal;
  public Boolean isPrimary;
  public Integer version;
  private BeneathClient client;

  public TableInstance() {
  }

  // TODO: this should probably follow the "Builder" pattern
  public static TableInstance make(BeneathClient client, Table table, CreateTableInstance adminData) {
    TableInstance instance = new TableInstance();
    instance.client = client;
    instance.table = table;
    instance.setAdminData(adminData);
    return instance;
  }

  // overloaded method to accommodate the
  // "CreateTableMutation.PrimaryTableInstance" type
  public static TableInstance make(BeneathClient client, Table table,
      CreateTableMutation.PrimaryTableInstance adminData) {
    TableInstance instance = new TableInstance();
    instance.client = client;
    instance.table = table;
    instance.setAdminData(adminData);
    return instance;
  }

  // overloaded method to accommodate the
  // "TableByOrganizationProjectAndNameQuery.PrimaryTableInstance" type
  public static TableInstance make(BeneathClient client, Table table,
      TableByOrganizationProjectAndNameQuery.PrimaryTableInstance adminData) {
    TableInstance instance = new TableInstance();
    instance.client = client;
    instance.table = table;
    instance.setAdminData(adminData);
    return instance;
  }

  public static TableInstance makeDry(BeneathClient client, Table table, Integer version, Boolean makePrimary) {
    TableInstance instance = new TableInstance();
    instance.client = client;
    instance.table = table;
    instance.instanceId = null;
    instance.version = version;
    instance.isPrimary = makePrimary;
    return instance;
  }

  private void setAdminData(CreateTableInstance adminData) {
    this.instanceId = UUID.fromString(adminData.tableInstanceID());
    this.version = adminData.version();
    this.isFinal = adminData.madeFinalOn() != null;
    this.isPrimary = adminData.madePrimaryOn() != null;
  }

  // overloaded method to accommodate the
  // "CreateTableMutation.PrimaryTableInstance" type
  private void setAdminData(CreateTableMutation.PrimaryTableInstance adminData) {
    this.instanceId = UUID.fromString(adminData.tableInstanceID());
    this.version = adminData.version();
    this.isFinal = adminData.madeFinalOn() != null;
    this.isPrimary = adminData.madePrimaryOn() != null;
  }

  // overloaded method to accommodate the
  // "TableByOrganizationProjectAndNameQuery.PrimaryTableInstance" type
  private void setAdminData(TableByOrganizationProjectAndNameQuery.PrimaryTableInstance adminData) {
    this.instanceId = UUID.fromString(adminData.tableInstanceID());
    this.version = adminData.version();
    this.isFinal = adminData.madeFinalOn() != null;
    this.isPrimary = adminData.madePrimaryOn() != null;
  }

  // READING RECORDS

  /**
   * Queries a sorted index of the records written to the table. The index
   * contains the newest record for each record key (see the table's schema for
   * the key). Returns a cursor for paging through the index.
   * 
   * @throws Exception
   */
  public Cursor queryIndex() throws Exception {
    return queryIndex("");
  }

  /**
   * Queries a sorted index of the records written to the table. The index
   * contains the newest record for each record key (see the table's schema for
   * the key). Returns a cursor for paging through the index.
   * 
   * Args: filter (str): A filter to apply to the index. Filters allow you to
   * quickly find specific record(s) in the index based on the record key. For
   * details on the filter syntax, see
   * https://about.beneath.dev/docs/reading-writing-data/index-filters/.
   */
  public Cursor queryIndex(String filter) throws Exception {
    // handle dry case
    if (this.instanceId == null) {
      throw new Exception("cannot query a dry instance");
    }
    QueryIndexResponse response = this.client.connection.queryIndex(this.instanceId, filter);
    assert response.getReplayCursorsCount() <= 1 && response.getChangeCursorsCount() <= 1;
    ByteString replay = response.getReplayCursorsCount() > 0 ? response.getReplayCursors(0) : null;
    ByteString changes = response.getChangeCursorsCount() > 0 ? response.getChangeCursors(0) : null;
    return new Cursor(this.client.connection, this.table.schema, replay, changes);
  }
}