"use client";

import {
  PinLeftIcon,
  PinRightIcon,
  HomeIcon,
  DashboardIcon,
  BarChartIcon,
  ChatBubbleIcon,
} from "@radix-ui/react-icons";
import { useState } from "react";
import Link from "next/link";

function Sidebar() {
  const [isCollapsed, setIsCollapsed] = useState(false);

  const handleSidebarToggle = () => {
    // console.log(`Inverting ${isCollapsed}`);
    setIsCollapsed(!isCollapsed);
  };

  return (
    <div className="h-full flex flex-col justify-between">
      <div className="flex flex-col gap-12">
        <div className="flex justify-between gap-2 px-5 pt-4 items-center">
          <div className="flex gap-2 justify-center items-center">
            <div className="text-xs">nlQ.ico</div>
            {!isCollapsed && <h1 className="text-2xl">nlQuery</h1>}
          </div>
          {!isCollapsed && (
            <button onClick={handleSidebarToggle}>
              <PinLeftIcon height={20} width={15} />
            </button>
          )}
        </div>
        <div className="flex flex-col gap-2 text-lg px-4 text-gray-600 ">
          <Link href="/">
            <div className="flex gap-4 items-center rounded-md transition-colors hover:text-gray-200 hover:bg-gray-700 px-4 py-2">
              <HomeIcon height={24} width={24} />
              {!isCollapsed && <p>Home</p>}
            </div>
          </Link>
          <Link href={"/charts"}>
            <div className="flex gap-4 items-center rounded-md transition-colors hover:text-gray-200 hover:bg-gray-700 px-4 py-2">
              <BarChartIcon height={24} width={24} />
              {!isCollapsed && <p>Charts</p>}
            </div>
          </Link>
          <Link href={"/dashboards"}>
            <div className="flex gap-4 items-center rounded-md transition-colors hover:text-gray-200 hover:bg-gray-700 px-4 py-2">
              <DashboardIcon height={24} width={24} />
              {!isCollapsed && <p>Dashboards</p>}
            </div>
          </Link>
          <Link href={"/support"}>
            <div className="flex gap-4 items-center rounded-md transition-colors hover:text-gray-200 hover:bg-gray-700 px-4 py-2">
              <ChatBubbleIcon height={24} width={24} />
              {!isCollapsed && <p>Support</p>}
            </div>
          </Link>
        </div>
      </div>

      <div className="flex flex-col items-center gap-2">
        {isCollapsed && (
          <button
            onClick={handleSidebarToggle}
            className="rounded-full hover:bg-gray-700 w-12 h-12 border-2 bg-gray-200 hover:stroke-gray-200"
          >
            <PinRightIcon
              height={20}
              width={15}
              className="mx-auto my-auto stroke-gray-600 hover:stroke-gray-200 transition-colors"
            />
          </button>
        )}
        {!isCollapsed && (
          <p className="w-48 p-2 text-center text-gray-600">
            Try nlQuery for your business today
          </p>
        )}
      </div>
    </div>
  );
}

export default Sidebar;
