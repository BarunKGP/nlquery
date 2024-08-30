"use client";

import { useRef, useState } from "react";
import { Button } from "@/components/ui/button";
import { ArrowRightIcon, PinTopIcon } from "@radix-ui/react-icons";
import { DropdownMenuRadioGroupDemo } from "@/components/component/radio-group";
import Papa from "papaparse";
import Plot from "react-plotly.js";
import RechartsComponent from "@/components/component/recharts-component";
import { useRouter } from "next/navigation";

const DATA = [
  { date: "2021-02-12", revenue: 1234.12 },
  { date: "2021-03-12", revenue: 1401.44 },
  { date: "2021-04-12", revenue: 988.43 },
];

function SearchArea() {
  const [contents, setContents] = useState("");
  const [files, setFiles] = useState("");
  const [llmData, setLlmData] = useState([]);
  const fileRef = useRef(null);
  const router = useRouter();

  const getLLMData = async (files, ms) => {
    setTimeout(() => console.log(`Waiting for ${ms / 1000} seconds`), ms);
  };

  const handleGo = async (e) => {
    // Not a form rn, so commenting out
    e.preventDefault();

    // We should not allow to submit if there are no contents
    // Gray out the Go button until there is content
    if (!contents) return;
    router.prefetch(`/recharts?prompt=${contents}`);
    if (!files) {
      console.log("No file attached");
      //! Remove after testing
      router.replace(`/recharts?prompt=${contents}`);
    }

    if (files) {
      console.log(files);
      router.push(`/recharts?prompt=${contents}`);
      // Get the Recharts code from LLM
      await getLLMData(files, 100);
      setLlmData(DATA);

      //* navigate to new page to display chart
      // TODO: Add animation in transition

      //* CSV parsing
      //? Should we send parsed CSV to backend?
      // Papa.parse(files, {
      //   header: true,
      //   dynamicTyping: true,
      //   complete: function (results) {
      //     const dailyRevenue = calculateDailyRevenue(results.data);
      //     const { dates, revenues } = prepareDataForPlotly(dailyRevenue);
      //     setLlmData({ x: dates, y: revenues });

      //     console.log("Plotly data has been set");
      //     // createBarGraph(dates, revenues);
      //   },
      // });
    }

    // Cleanup the previous prompts
    setContents("");
    setFiles(null);
  };

  const handleFileChange = (e) => {
    const file = e.target.files[0];
    setFiles(file ? file.name : "");
    console.log(file);
  };

  return (
    <div className="grid w-full grid-rows-3 gap-4 rounded-lg">
      <div className="flex flex-col justify-center gap-2 ">
        <h1 className="pt-4 text-5xl tracking-tighter text-center text-gray-700">
          Charts. In plain English.
        </h1>
        <h2 className="text-center text-gray-700">
          Creating reports has never been easier
        </h2>
      </div>
      <div>
        <div className="flex flex-col justify-center mx-auto border-2 border-gray-300 rounded-md w-fit h-fit">
          <div
            id="search-text-area"
            className="px-6 py-2 md:min-w-[500px] sm:min-w-[500px]"
          >
            <textarea
              type="text"
              value={contents}
              onChange={(e) => {
                setContents(e.target.value);
              }}
              className="w-full appearance-none resize-none min-h-32 max-h-fit focus-visible:outline-none bg-stone-100"
              placeholder="Ask for visualization..."
            />
          </div>
          <div className="flex items-center justify-between w-full p-2">
            <div className="flex items-center gap-2">
              <input
                type="file"
                onChange={handleFileChange}
                className="hidden"
                id="file-attach"
                ref={fileRef}
              />
              <label htmlFor="file-attach">
                <Button
                  variant="ghost"
                  onClick={() => {
                    fileRef.current.click();
                  }}
                >
                  <div className="flex items-center justify-between h-full gap-1 text-gray-400 align-middle transition-colors stroke-gray-400 hover:stroke-800 hover:text-gray-800">
                    <PinTopIcon height={12} width={12} />
                    <span className="text-sm"> Attach File </span>
                  </div>
                </Button>
              </label>
              {files && (
                <span className="text-sm font-bold text-cyan-700 hover:text-cyan-600">
                  {files}
                </span>
              )}
            </div>
            <Button
              variant="outline"
              type="submit"
              onClick={(e) => handleGo(e)}
              className="w-12 h-12 rounded-full"
            >
              <div className="flex items-center gap-2 tracking-wide text-gray-400 transition-colors stroke-gray-400 hover:stroke-gray-900 hover:text-gray-900">
                {/* <span className="text-sm">Go</span> */}
                <ArrowRightIcon height={20} width={20} />
              </div>
            </Button>
          </div>
        </div>
      </div>
      <div id="displayComponent" className="last:mb-32">
        {llmData.length > 0 && (
          <div className="flex justify-between">
            <RechartsComponent
              width={800}
              height={500}
              // data={llmData}
              data={DATA}
              xDataKey="date"
            />
            <Button className="mt-5">Customize further</Button>
          </div>
        )}
      </div>
    </div>
  );
}



export default SearchArea;
