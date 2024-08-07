"use client";
import React from "react";
import Image from "next/image";
import ConnectorSvg from "@/components/component/connector-svg";
import connectorIconUrl from "../../public/reshot-icon-connector.svg?url";
import queryIconUrl from "../../public/reshot-icon-query.svg?url";
import dashboardIconUrl from "../../public/reshot-icon-dashboard.svg?url";
import shareIconUrl from "../../public/reshot-icon-share.svg?url";

function FeatureListItem({ connectorKey }) {
  const classNames = "flex justify-center h-32";
  if (connectorKey === "connector") {
    return (
      <div className="grid gap-2 text-themetext text-center">
        {/* <ConnectorIcon /> */}
        <div className={classNames}>
          <Image
            src={connectorIconUrl}
            alt="connector-icon-svg"
            height={96}
            width={96}
          />
        </div>
        <h2 className="text-xl text-center">Connect to your data sources</h2>
        <p className="text-md">
          We only support CSVs at the moment. SQL support coming soon
        </p>
      </div>
    );
  } else if (connectorKey === "dashboard") {
    return (
      <div className="grid gap-2 text-themetext text-center">
        {/* <ConnectorIcon /> */}
        <div className={classNames}>
          <Image
            src={dashboardIconUrl}
            alt="dashboard-icon-svg"
            height={80}
            width={80}
          />
        </div>
        <h2 className="text-xl text-center">Generate Dashboard</h2>
        <p className="text-md">
          Drag and drop tiles with custom interactive charts generated from
          querying data sources
        </p>
      </div>
    );
  } else if (connectorKey === "query") {
    return (
      <div className="grid gap-2 text-themetext text-center">
        {/* <ConnectorIcon /> */}
        <div className={classNames}>
          <Image
            src={queryIconUrl}
            alt="query-icon-svg"
            height={80}
            width={80}
          />
        </div>
        <h2 className="text-xl text-center">Queries in Natural Language</h2>
        <p className="text-md italic">
          “Generate a bar chart showing revenue of different products sold last
          month”
        </p>
      </div>
    );
  } else if (connectorKey === "share") {
    return (
      <div className="grid gap-2 text-themetext text-center">
        {/* <ConnectorIcon /> */}
        <div className={classNames}>
          <Image
            src={shareIconUrl}
            alt="share-icon-svg"
            height={80}
            width={80}
          />
        </div>
        <h2 className="text-xl text-center">Share Insights Instantly</h2>
        <p className="text-md">
          Data refreshes at periodic intervals to enable realtime decision
          making
        </p>
      </div>
    );
  } else {
    console.log("Invalid key");
    return <div>Invalid key</div>;
  }
}

function FeatureList({ listKeys }) {
  return (
    <div className="w-4/5 mx-auto grid grid-cols-4 gap-4 justify-center p-10">
      {listKeys.map((key, index) => {
        return <FeatureListItem connectorKey={key} key={index} />;
      })}
    </div>
  );
}

export default FeatureList;
