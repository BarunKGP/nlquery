import React from "react";
import Image from "next/image";
import NlQueryLogo from "../../public/nlquery-logo.png";

function NlQueryBanner() {
  return (
    <div>
      <div className="m-2 gradient-border-parent">
        <div className="flex items-center justify-center gap-1 px-2 gradient-border-child">
          <Image src={NlQueryLogo} width={60} height={40} />
          <span
            className="font-semibold tracking-tighter stippling text-text"
            data-text="nlQuery"
          >
            nlQuery
          </span>
        </div>
      </div>
    </div>
  );
}

export default NlQueryBanner;
