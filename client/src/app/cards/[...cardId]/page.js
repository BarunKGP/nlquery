import React from "react";
import SearchArea from "./SearchArea";
import Sidebar from "@/components/Sidebar";

function Page() {
  return (
    <div className="min-h-screen flex justify-between">
      <div className="border-2 border-bg ">
        <Sidebar />
      </div>
      <div className="w-full md:w-1/2 grid gap-4 grid-rows-7 rounded-lg">
        <div className="flex flex-col justify-center gap-2">
          <h1 className="text-5xl tracking-tighter text-center row-span-1 pt-4 text-gray-700">
            Charts. In plain English.
          </h1>
          <h2 className="text-center text-gray-700">
            Creating reports has never been easier
          </h2>
        </div>
        <div className="row-span-2">
          {/* <div className="items-stretch p-2 ml-auto">
            <Button variant="default" className="p-2">
              Add data
            </Button>
          </div> */}
          <SearchArea />
        </div>
        {/* <DisplayComponent /> */}
        <div id="displayComponent"></div>
      </div>
      <div>Right Sidebar</div>
    </div>
  );
}

export default Page;
