"use client";

import { useEffect, useState } from "react";
import { useSession } from "next-auth/react";

import CollapsibleSidebar from "./sidebar";
import QuickActions from "./quick-actions";
import DashboardContents from "./dashboard-contents";
import { Button } from "@/components/ui/button";

import { Send } from "lucide-react";

function Page() {
  const { data: session, status } = useSession();
  const [name, setName] = useState("");

  useEffect(() => {
    if (session && status === "authenticated") {
      const name = session.user.name.split(" ")[0];
      console.log(`User id: ${session.user.id}`);
      setName(name);
    }
  }, [status]);

  return (
    <div className="w-screen min-h-screen bg-background">
      <div className="flex h-screen">
        <CollapsibleSidebar />
        <div className="grid items-center w-full gap-4 justify-items-center grid-rows-7">
          <div className="grid w-full grid-cols-2 px-4 justify-items-end">
            <h1 className="space-x-1 text-3xl tracking-tight text-center text-text">
              <span>Good evening, </span>
              <span className="italic">{name}</span>
            </h1>
            <div>
              <Button className="space-x-2 font-semibold bg-blue-600 text-text-inverse">
                <Send className="stroke-text-inverse" />
                <span>Invite</span>
              </Button>
            </div>
          </div>
          <QuickActions className="row-span-2" />
          <div className="w-full row-span-3 px-4 space-y-2">
            <DashboardContents />
            {/* <table className="w-full text-center text-text">
              <tr className="text-sm text-text-muted">
                <th>NAME</th>
                <th>TYPE</th>
                <th>CREATED AT</th>
              </tr>
              <tr>
                <td>Test file</td>
                <td>Chart</td>
                <td>08-31-2024 11:56 PM</td>
              </tr>
              <tr>
                <td>Test file</td>
                <td>Chart</td>
                <td>08-31-2024 11:56 PM</td>
              </tr>
            </table> */}
          </div>
          <div className="flex justify-center w-full row-span-1">
            <Footer />
          </div>
        </div>
      </div>
    </div>
  );
}

const Footer = () => {
  return (
    <div className="flex gap-20 p-2 mx-4 mb-4 border-2 rounded-full opacity-70 w-fit border-text-muted">
      <div className="text-xs text-center text-text-muted -tracking-tight">
        Terms of Service
      </div>
      <div className="text-xs text-center text-text-muted -tracking-tight">
        Privacy Policy
      </div>
      <div className="text-xs text-center text-text-muted -tracking-tight">
        Support
      </div>
      <div className="text-xs text-center text-text-muted -tracking-tight">
        Pricing
      </div>
      <div className="text-xs text-center text-text-muted -tracking-tight">
        Enterprise
      </div>
      <div className="text-xs text-center text-text-muted -tracking-tight">
        2024 &copy; nlQuery
      </div>
    </div>
  );
};

export default Page;
