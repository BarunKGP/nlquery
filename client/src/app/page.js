import React from "react";
import { Button } from "@/components/ui/button";

export default function Home() {
  return (
    <div>
      <div className="flex justify-center gap-8 align-middle curved-edge text-stone-200">
        <div className="flex flex-col justify-center gap-6 p-4 h-3/4 max-w-[800px]">
          <h1 className="text-5xl font-bold tracking-tighter text-themetext">
            Build interactive visualizations by talking to your data
            {/* <span className="pr-2 italic">nl</span>
              <span className="text-themetext">Query</span> */}
          </h1>
          <h2 className="text-lg tracking-tight text-textcolor">
            As engineers, we wanted to make it incredibly simple to generate
            insightful visualizations to empower decision-makers. We built
            nlQuery, a powerful visualization engine which interprets natural
            language queries to generate beautiful, customizable visualizations
            suitable for any business use case
          </h2>
          <div className="flex gap-2">
            {/* <Button className="text-lg">Start Building</Button> */}
            <Button variant="secondary" className="text-lg text-dgreen1">
              Get on Waitlist
            </Button>
          </div>
        </div>
        <div>
          <div className="w-[500px] border-2 h-96 border-stone-200">
            Placeholder
          </div>
        </div>
      </div>
      <div className="min-h-[40vh] text-center">Content Placeholder</div>
    </div>
  );
}
