import React from "react";
import {
  LineChart,
  Line,
  CartesianGrid,
  XAxis,
  YAxis,
  Tooltip,
} from "recharts";

const customTicks = ({ x, y, payload }) => {
  dt = new Date(payload.value);
  return `${dt.getDay()} ${dt.getMonth() + 1} ${dt.getFullYear()}`;
};

function Recharts() {
  return (
    <div>
      <LineChart
        width={700}
        height={500}
        data={[
          { date: "2021-02-12", revenue: 1234.12 },
          { date: "2021-03-12", revenue: 1401.44 },
          { date: "2021-04-12", revenue: 988.43 },
        ]}
      >
        <Line type="monotone" dataKey="revenue" stroke="rgb(179, 50, 91)" />
        <CartesianGrid stroke="#ccc" strokeDasharray="5 5" />
        <XAxis dataKey="date" />
        <YAxis />
        <Tooltip tick={customTicks} />
      </LineChart>
    </div>
  );
}

export default Recharts;
