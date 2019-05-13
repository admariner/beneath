import { gql } from "apollo-server";
import { GraphQLResolveInfo } from "graphql";
import _ from "lodash";

import { Project } from "../entities/Project";
import { canEditProject, canReadProject } from "../lib/guards";
import { IApolloContext } from "../types";

export const typeDefs = gql`
  extend type Query {
    project(name: String, projectId: ID): Project
  }

  type Project {
    projectId: ID!
    name: String
    displayName: String
    site: String
    description: String
    createdOn: Date
    updatedOn: Date
    users: [User]
    canEdit: Boolean
  }
`;

export const resolvers = {
  Query: {
    project: async (root: any, args: any, ctx: IApolloContext, info: GraphQLResolveInfo) => {
      const project = await Project.findOne(args, { relations: ["users"] });

      await canReadProject(ctx, project.projectId);

      const meUserId = _.get(ctx, "user.key.userId");
      const canEdit = project.users.some((user) => user.userId === meUserId);

      return {
        projectId: project.projectId,
        name: project.name,
        displayName: project.displayName,
        site: project.site,
        description: project.description,
        createdOn: project.createdOn,
        updatedOn: project.updatedOn,
        users: project.users,
        canEdit,
      };
    },
  },
};
