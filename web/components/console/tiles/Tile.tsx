import Link from "next/link";
import React, { FC } from "react";
import { Grid, makeStyles, Paper } from "@material-ui/core";

import clsx from "clsx";

const useStyles = makeStyles(() => ({
  paper: {
    height: "100%",
  },
  paperLink: {
    cursor: "pointer",
  },
  unstyledA: {
    color: "inherit",
    textDecoration: "inherit",
  }
}));

// tslint:disable-next-line: no-empty-interface
export interface TileProps {
  shape?: "normal" | "wide" | "dense";
  href?: string;
  as?: string;
  nopaper?: boolean;
  responsive?: boolean;
  styleClass?: string;
}

export const Tile: FC<TileProps> = ({ shape, href, as, nopaper, responsive, styleClass, children }) => {
  if (!shape) {
    shape = "normal";
  }
  if (responsive === undefined) {
    responsive = true;
  }

  const classes = useStyles();

  let tile;
  if (nopaper) {
    tile = <div className={clsx(classes.paper, href && classes.paperLink)}>{children}</div>;
  } else {
    tile = <Paper className={clsx(classes.paper, href && classes.paperLink)}>{children}</Paper>;
  }

  if (href) {
    if (as) {
      tile = <Link href={href} as={as}>{tile}</Link>;
    } else {
      tile = <a className={classes.unstyledA} href={href}>{tile}</a>;
    }
  }
  return (
    <>
      {shape === "dense" && <Grid item>{tile}</Grid>}
      {shape === "normal" && responsive && (
        <Grid item xs={12} sm={6} md={4} lg={3} className={styleClass}>
          {tile}
        </Grid>
      )}
      {shape === "normal" && !responsive && (
        <>
          <Grid item className={styleClass}>
            {tile}
          </Grid>
        </>
      )}
      {shape === "wide" && (
        <Grid item xs={12} sm={12} md={8} lg={6}>
          {tile}
        </Grid>
      )}
    </>
  );
};

export default Tile;
