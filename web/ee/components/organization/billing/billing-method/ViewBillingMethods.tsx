import {useMutation, useQuery } from "@apollo/client";
import _ from "lodash";
import { Button, Grid, makeStyles, Table, TableBody, TableCell, TableHead, TableRow } from "@material-ui/core";
import { MoreVert } from "@material-ui/icons";
import React, { FC } from "react";

import { OrganizationByName_organizationByName_PrivateOrganization } from "apollo/types/OrganizationByName";
import ContentContainer, { CallToAction } from "components/ContentContainer";
import DropdownButton from "components/DropdownButton";
import { QUERY_BILLING_METHODS } from "ee/apollo/queries/billingMethod";
import { BillingMethods, BillingMethodsVariables } from "ee/apollo/types/BillingMethods";
import { BillingInfo_billingInfo, BillingInfo_billingInfo_billingMethod } from "ee/apollo/types/BillingInfo";
import { UpdateBillingMethod, UpdateBillingMethodVariables } from "ee/apollo/types/UpdateBillingMethod";
import { UPDATE_BILLING_METHOD } from "ee/apollo/queries/billingInfo";
import {ANARCHISM_DRIVER, STRIPECARD_DRIVER, STRIPEWIRE_DRIVER} from "ee/lib/billing";

const useStyles = makeStyles((theme) => ({
  container: {
    overflowX: "auto",
  },
  button: {
    marginTop: theme.spacing(4),
  },
  dropdownButton: {
    backgroundColor: theme.palette.background.paper,
    "&:hover": {
      backgroundColor: theme.palette.secondary.main,
    },
    boxShadow: "none",
    color: theme.palette.common.white,
  }
}));

export interface BillingMethodsProps {
  organization: OrganizationByName_organizationByName_PrivateOrganization;
  billingInfo: BillingInfo_billingInfo;
  addCard: (value: boolean) => void;
}

const ViewBillingMethods: FC<BillingMethodsProps> = ({ organization, billingInfo, addCard }) => {
  const classes = useStyles();
  const { loading, error, data } = useQuery<BillingMethods, BillingMethodsVariables>(QUERY_BILLING_METHODS, {
    context: { ee: true },
    variables: { organizationID: organization.organizationID },
  });

  const [updateBillingMethod] = useMutation<UpdateBillingMethod, UpdateBillingMethodVariables>(UPDATE_BILLING_METHOD, {
    context: { ee: true }
  });

  if (error) {
    return <p>Error: {JSON.stringify(error)}</p>;
  }

  if (!data) {
    return <></>;
  }

  if (data.billingMethods.length === 0) {
    const cta: CallToAction = {
      message: `You have no billing methods on file`,
      buttons: [{ label: "Add a credit card", onClick: () => addCard(true) }]
    };
    return (
      <ContentContainer callToAction={cta} />
    );
  }

  const formatDriver = (driver: string) => {
    if (driver === STRIPECARD_DRIVER) return "Card";
    if (driver === STRIPEWIRE_DRIVER) return "Wire";
    if (driver === ANARCHISM_DRIVER) return "Anarchism";
  };

  const formatDetails = (driver: string, driverPayload: string) => {
    if (driver === STRIPECARD_DRIVER) {
      const payload = JSON.parse(driverPayload);
      const brand = payload.brand.toString();
      const last4 = payload.last4.toString();
      const expMonth = payload.expMonth.toString();
      const expYear = payload.expYear.toString();
      return `${brand.toUpperCase()} ${last4}, Exp: ${expMonth}/${expYear.substring(2,4)}`;
    }
    if (driver === STRIPEWIRE_DRIVER) {
      const payload = JSON.parse(driverPayload);
      return `Invoices sent to ${payload.email_address}`;
    }
  };

  const getBillingMethodActions = (billingMethod: BillingInfo_billingInfo_billingMethod) => [
    { label: "Set to active", onClick: () => updateBillingMethod({
      variables: {
        organizationID: organization.organizationID,
        billingMethodID: billingMethod.billingMethodID
      }})
    },
    // TODO: enable deleting billing methods
    // { label: "Delete billing method", onClick: () => deleteBillingMethod() }
  ];

  return (
    <>
      <Table className={classes.container}>
        <TableHead>
          <TableRow>
            <TableCell>Type</TableCell>
            <TableCell>Details</TableCell>
            <TableCell>Active</TableCell>
            <TableCell></TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {data.billingMethods.map((billingMethod) => (
            <TableRow
              key={billingMethod.billingMethodID}
            >
              <TableCell>{formatDriver(billingMethod.paymentsDriver)}</TableCell>
              <TableCell>{formatDetails(billingMethod.paymentsDriver, billingMethod.driverPayload)}</TableCell>
              <TableCell>{billingMethod.billingMethodID === billingInfo.billingMethod?.billingMethodID ? "Yes" : "No"}</TableCell>
              <TableCell>
                <DropdownButton
                  variant="contained"
                  margin="dense"
                  actions={getBillingMethodActions(billingMethod)}
                  className={classes.dropdownButton}
                >
                  <MoreVert />
                </DropdownButton>
                </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
      <Button variant="contained" onClick={() => addCard(true)} className={classes.button}>Add card</Button>
    </>
  );
};

export default ViewBillingMethods;