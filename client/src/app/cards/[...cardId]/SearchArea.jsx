"use client";

import { useRef, useState } from "react";
import { Button } from "@/components/ui/button";
import { ArrowRightIcon, PinTopIcon } from "@radix-ui/react-icons";
import { DropdownMenuRadioGroupDemo } from "@/components/component/radio-group";

function SearchArea() {
  const [contents, setContents] = useState("");
  const [files, setFiles] = useState("");
  const fileRef = useRef(null);

  const handleGo = (e) => {
    // e.preventDefault();

    if (!contents) return;
    if (!files) {
      console.log("No file attached");
    }

    console.log(contents);
    if (files) {
      console.log(files);
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
    <div className="border-2 border-gray-300 rounded-md flex flex-col justify-between ">
      <div id="search-text-area" className="px-6 py-2">
        <textarea
          type="text"
          value={contents}
          onChange={(e) => {
            setContents(e.target.value);
          }}
          className="w-full min-h-44 max-h-fit resize-none focus-visible:outline-none appearance-none"
          placeholder="Ask for visualization..."
        />
      </div>
      <div className="flex w-full justify-between p-2 items-center">
        <div className="flex gap-2 items-center">
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
              <div
                className="h-full flex justify-between gap-1 items-center
              stroke-gray-400 text-gray-400 align-middle hover:stroke-800 
              transition-colors hover:text-gray-800"
              >
                <PinTopIcon height={12} width={12} />
                <span className="text-sm"> Attach File </span>
              </div>
            </Button>
          </label>
          {files && (
            <span className="text-cyan-700 text-sm font-bold hover:text-cyan-600">
              {files}
            </span>
          )}
        </div>
        <Button
          variant="outline"
          type="submit"
          onClick={(e) => handleGo(e)}
          className="rounded-full h-12 w-12"
        >
          <div className="flex gap-2 items-center text-gray-400 stroke-gray-400 tracking-wide hover:stroke-gray-900 hover:text-gray-900 transition-colors">
            {/* <span className="text-sm">Go</span> */}
            <ArrowRightIcon height={20} width={20} />
          </div>
        </Button>
      </div>
    </div>
  );
}

export default SearchArea;
