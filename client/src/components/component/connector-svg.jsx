import * as React from "react";

const ConnectorSvg = (props) => {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      xmlSpace="preserve"
      width={48}
      height={48}
      style={{
        shapeRendering: "geometricPrecision",
        textRendering: "geometricPrecision",
        imageRendering: "optimizeQuality",
        fillRule: "evenodd",
        clipRule: "evenodd",
      }}
      {...props}
    >
      <defs>
        <style>{".fil0{fill:#212121;fill-rule:nonzero}"}</style>
      </defs>
      <g id="Layer_x0020_1">
        <path
          d="M891.565 1498.5c17.052 4.556 34.57-5.574 39.125-22.625 4.556-17.052-5.573-34.57-22.625-39.125-93.691-25.105-173.079-80.065-229.016-152.965-55.969-72.942-88.53-163.898-88.53-260.945 0-118.324 47.958-225.444 125.495-302.98 77.536-77.537 184.656-125.495 302.98-125.495 17.673 0 32.001-14.328 32.001-32.001s-14.328-32.001-32-32.001c-135.99 0-259.11 55.123-348.232 144.246-89.122 89.121-144.246 212.242-144.246 348.231 0 111.366 37.434 215.834 101.782 299.694 64.38 83.9 155.645 147.128 263.266 175.965z"
          className="fil0"
        />
        <path
          d="M1051 278.997c0-17.673-14.328-32-32-32-17.674 0-32.002 14.327-32.002 32v283.37c0 17.673 14.328 32 32.001 32s32.001-14.327 32.001-32v-283.37zM1151.43 540.495c-17.052-4.556-34.57 5.574-39.125 22.625-4.556 17.052 5.573 34.57 22.625 39.125 93.69 25.105 173.078 80.065 229.016 152.966 55.969 72.942 88.53 163.898 88.53 260.945 0 118.323-47.96 225.443-125.495 302.98-77.536 77.536-184.657 125.494-302.98 125.494-17.673 0-32.001 14.328-32.001 32.001s14.328 32.001 32 32.001c135.99 0 259.11-55.123 348.232-144.246 89.12-89.121 144.246-212.241 144.246-348.229 0-111.366-37.434-215.834-101.782-299.694-64.38-83.903-155.645-147.129-263.266-175.967z"
          className="fil0"
        />
        <path
          d="M992.002 1760c0 17.673 14.328 32 32 32 17.674 0 32.002-14.327 32.002-32v-283.37c0-17.673-14.328-32-32.001-32s-32.001 14.327-32.001 32V1760z"
          className="fil0"
        />
      </g>
      <path
        d="M0 0h2048v2048H0z"
        style={{
          fill: "none",
        }}
      />
    </svg>
  );
};

export default ConnectorSvg;
