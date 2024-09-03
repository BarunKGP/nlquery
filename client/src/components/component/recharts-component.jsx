"use client";

import React, { ReactNode } from "react";
import { Bar, BarChart, Legend, Rectangle, XAxis, YAxis } from "recharts";

// interface RechartsComponentProps {
//   children: ReactNode;
//   width: number;
//   height: number;
//   data: any;
//   xDataKey: string;
// }

export default function RechartsComponent({
  children,
  width,
  height,
  data,
  xDataKey,
}) {
  return (
    <div>
      <BarChart width={width} height={height} data={data}>
        <XAxis dataKey={xDataKey} stroke="rgb(179,50,91)" />
        <YAxis />
        <Legend
          width={100}
          wrapperStyle={{
            top: 40,
            right: 20,
            backgroundColor: "#f5f5f5",
            border: "1px solid #d5d5d5",
            borderRadius: 3,
            color: "rgb(67, 56, 202)",
          }}
        />
        <Bar
          dataKey="revenue"
          fill="rgb(212, 89, 134)"
          activeBar={<Rectangle fill="orange" stroke="pink" />}
          legendType="circle"
        />
      </BarChart>
    </div>
  );
}
