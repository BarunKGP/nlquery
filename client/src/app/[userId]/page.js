import React from "react";
import SearchArea from "./SearchArea";

function Page() {
  return (
    <div className="min-h-screen flex justify-center gap-12 p-4">
      <div>Left sidebar</div>
      <div className="w-1/2 grid gap-4 grid-rows-7 border-2 border-headerbg p-4 shadow-lg rounded-lg">
        <h1 className="text-5xl tracking-tighter text-center row-span-1">
          Build card
        </h1>
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
