"use client";
// import { SearchBar } from "@/components/component/search-bar";
import { useRef, useState } from "react";
import { Button } from "@/components/ui/button";
import { ArrowRightIcon, PinTopIcon } from "@radix-ui/react-icons";
import { cn } from "@/lib/utils";
import { DropdownMenuRadioGroupDemo } from "@/components/component/radio-group";
// function submitData({data}) {

// }

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

  const handleFileButtonClick = () => {
    fileRef.current.click();
  };

  return (
    <div className="border-2 border-gray-300 rounded-md h-3/5 flex flex-col justify-between">
      <div id="search-text-area" className="px-6 py-2">
        <textarea
          type="text"
          value={contents}
          onChange={(e) => {
            setContents(e.target.value);
          }}
          className="w-full h-full resize-none focus-visible:outline-none appearance-none"
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
            <span className="text-cyan-700 text-sm font-semibold hover:text-cyan-600">
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

function OldSearchArea() {
  const [contents, setContents] = useState("");
  const [attachment, setAttachment] = useState(null);

  const handleSubmit = (e) => {
    e.preventDefault();
    // Do error handling.
    // Error can be a simple state object

    console.log("Entering file event handler", e);

    if (!contents) return;
    if (contents.length > 0) {
      console.log(contents);
    }
  };
  return (
    <form className="h-full">
      <div
        className="p-2 border-2 border-gray-200 rounded-md h-2/3 
      disabled:cursor-not-allowed disabled:opacity-50 focus-visible:ring-1 
      focus-visible:border-gray-400 transition-colors flex flex-col justify-between"
      >
        <div>
          <input
            type="text"
            value={contents}
            onChange={(e) => setContents(e.target.value)}
          />
          <textarea
            className="px-4 rounded-md w-full resize-none
                  focus-visible:outline-none appearance-none tracking-wide text-wrap"
            placeholder="Ask for visualization"
          ></textarea>
        </div>
        <div className="flex justify-between">
          <div className="flex justify-evenly w-full">
            <div>
              <input
                type="file"
                hidden
                onChange={(e) => setAttachment(e.target.files[0])}
              />
              <label>
                <Button variant="ghost">
                  <div
                    className="h-full flex tracking-wide justify-around gap-2 stroke-gray-400 text-gray-400
                          align-middle hover:stroke-800 transition-colors hover:text-gray-800"
                  >
                    <PinTopIcon height={16} width={14} />
                    <span> Attach File </span>
                  </div>
                </Button>
              </label>
              {attachment ? <label>{attachment}</label> : <></>}
            </div>
            <div></div>
          </div>
          <Button
            variant="outline"
            type="submit"
            onClick={(e) => handleSubmit(e)}
            className="rounded-3xl flex gap-2 align-middle text-gray-400 stroke-gray-400 tracking-wide hover:stroke-gray-900 hover:text-gray-900 hover:bg-themetext transition-colors"
            // className={cn(
            //   "rounded-3xl flex gap-2 align-middle text-gray-400 stroke-gray-400 tracking-wide hover:stroke-gray-800 hover:text-gray-800 hover:bg-themetext transition-colors",
            //   contents.length > 0 ? "bg-amber-400" : "bg-transparent"
            // )}
          >
            <span>Go </span>
            <ArrowRightIcon height={17} width={20} />
          </Button>
        </div>
      </div>
    </form>
  );
}

export default SearchArea;
