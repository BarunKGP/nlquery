import { auth } from "@/auth";
import SearchArea from "./SearchArea";
import Sidebar from "@/components/Sidebar";
import Image from "next/image";

async function Page() {
  const session = await auth();
  if (session == null) {
    return <p>Not logged in</p>;
  } else {
    return (
      <div className="flex justify-between w-full bg-stone-100">
        <div className="border-2 border-bg ">
          <Sidebar />
        </div>
        <SearchArea />
        <div>{`Client side: ${JSON.stringify(session, null, 2)}`}</div>
      </div>
    );
  }
}

export default Page;
