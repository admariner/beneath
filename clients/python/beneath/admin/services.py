from beneath.connection import Connection
from beneath.utils import format_entity_name


class Services:
    def __init__(self, conn: Connection):
        self.conn = conn

    async def find_by_organization_project_and_name(
        self, organization_name, project_name, service_name
    ):
        result = await self.conn.query_control(
            variables={
                "organizationName": format_entity_name(organization_name),
                "projectName": format_entity_name(project_name),
                "serviceName": format_entity_name(service_name),
            },
            query="""
                query ServiceByOrganizationProjectAndName(
                    $organizationName: String!
                    $projectName: String!
                    $serviceName: String!
                ) {
                    serviceByOrganizationProjectAndName(
                        organizationName: $organizationName
                        projectName: $projectName
                        serviceName: $serviceName
                    ) {
                        serviceID
                        name
                        description
                        sourceURL
                        readQuota
                        writeQuota
                        scanQuota
                    }
                }
            """,
        )
        return result["serviceByOrganizationProjectAndName"]

    async def stage(
        self,
        organization_name,
        project_name,
        service_name,
        description=None,
        source_url=None,
        read_quota_bytes=None,
        write_quota_bytes=None,
        scan_quota_bytes=None,
    ):
        result = await self.conn.query_control(
            variables={
                "organizationName": format_entity_name(organization_name),
                "projectName": format_entity_name(project_name),
                "serviceName": format_entity_name(service_name),
                "description": description,
                "sourceURL": source_url,
                "readQuota": read_quota_bytes,
                "writeQuota": write_quota_bytes,
                "scanQuota": scan_quota_bytes,
            },
            query="""
                mutation StageService(
                    $organizationName: String!
                    $projectName: String!
                    $serviceName: String!
                    $description: String
                    $sourceURL: String
                    $readQuota: Int
                    $writeQuota: Int
                    $scanQuota: Int
                ) {
                    stageService(
                        organizationName: $organizationName
                        projectName: $projectName
                        serviceName: $serviceName
                        description: $description
                        sourceURL: $sourceURL
                        readQuota: $readQuota
                        writeQuota: $writeQuota
                        scanQuota: $scanQuota
                    ) {
                        serviceID
                        name
                        description
                        sourceURL
                        readQuota
                        writeQuota
                        scanQuota
                    }
                }
            """,
        )
        return result["stageService"]

    async def update_permissions_for_stream(self, service_id, stream_id, read, write):
        result = await self.conn.query_control(
            variables={
                "serviceID": service_id,
                "streamID": stream_id,
                "read": read,
                "write": write,
            },
            query="""
                mutation UpdateServicePermissions(
                    $serviceID: UUID!
                    $streamID: UUID!
                    $read: Boolean
                    $write: Boolean
                ) {
                    updateServiceStreamPermissions(
                        serviceID: $serviceID
                        streamID: $streamID
                        read: $read
                        write: $write
                    ) {
                        serviceID
                        streamID
                        read
                        write
                    }
                }
            """,
        )
        return result["updateServiceStreamPermissions"]

    async def delete(self, service_id):
        result = await self.conn.query_control(
            variables={
                "serviceID": service_id,
            },
            query="""
                mutation DeleteService($serviceID: UUID!) {
                    deleteService(serviceID: $serviceID)
                }
            """,
        )
        return result["deleteService"]

    async def issue_secret(self, service_id, description):
        result = await self.conn.query_control(
            variables={
                "serviceID": service_id,
                "description": description,
            },
            query="""
                mutation IssueServiceSecret($serviceID: UUID!, $description: String!) {
                    issueServiceSecret(serviceID: $serviceID, description: $description) {
                        token
                    }
                }
            """,
        )
        return result["issueServiceSecret"]

    async def list_secrets(self, service_id):
        result = await self.conn.query_control(
            variables={
                "serviceID": service_id,
            },
            query="""
                query SecretsForService($serviceID: UUID!) {
                    secretsForService(serviceID: $serviceID) {
                        secretID
                        description
                        prefix
                        createdOn
                        updatedOn
                    }
                }
            """,
        )
        return result["secretsForService"]
