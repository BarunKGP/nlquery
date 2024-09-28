import CollapsibleSidebar from "./sidebar";
import QuickActions from "./quick-actions";
import DashboardContents from "./dashboard-contents";
import CallBackend from "./call-backend";

import { auth } from "@/auth";
import { Button } from "@/components/ui/button";
import ModeSwitcher from "@/components/mode-switcher";
import NlqueryBanner from "@/components/nlquery-logo-banner";

import Link from "next/link";
import { Send } from "lucide-react";

async function Page() {
  const session = await auth();

  if (!session || !session.user)
    // TODO: Refactor into own page for protected routes
    return (
      <div>
        <Header />
        <div className="p-4 mx-auto my-auto border-2 rounded-lg mt-[200px] h-fit text-text w-fit bg-secondary border-primary">
          <div className="flex flex-col items-center justify-center gap-2">
            <NlqueryBanner />
            <p className="text-xl font-semibold tracking-tight">
              You must be signed in to proceed
            </p>
          </div>
          <div className="flex items-center justify-center gap-6 mt-12">
            <Link href={"/login"}>
              <Button>Log in</Button>
            </Link>
            <Link href={"/login"}>
              <Button variant="outline">Sign up</Button>
            </Link>
          </div>
        </div>
      </div>
    );

  const name = session.user.name.split(" ")[0];

  return (
    <div className="flex w-screen h-screen bg-background">
      <CollapsibleSidebar />
      <div className="flex w-full mt-20 ">
        {/* <Header message={`Good evening, ${name}`} /> */}
        <div className="grid items-center w-full gap-4 mt-4 justify-items-center grid-rows-7">
          {/*            
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
          */}
          <QuickActions />
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
            <CallBackend />
          </div>
          <div className="flex justify-center w-full row-span-1">
            <Footer />
          </div>
        </div>
      </div>
    </div>
  );
}

const Header = ({ message }) => {
  return (
    <div className="absolute flex items-center justify-between w-full h-20 px-6 py-4 border-b-2 border-orange-500 top-1">
      <div>Sidebar Button</div>
      <div className="font-semibold">{message}</div>
      <div className="flex items-center justify-center gap-6 px-2">
        <div>
          <Button className="space-x-2 font-semibold bg-blue-600 text-text-inverse">
            <Send className="stroke-text-inverse" />
            <span>Invite</span>
          </Button>
        </div>
        <div>
          <ModeSwitcher />
        </div>
      </div>
    </div>
  );
};

const Footer = () => {
  return (
    <div className="absolute flex items-center justify-center gap-20 p-2 mx-4 mb-4 border-2 rounded-full opacity-70 w-fit border-text-muted bottom-1">
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
