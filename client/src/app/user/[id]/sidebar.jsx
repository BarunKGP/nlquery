"use client";

import { useSession } from "next-auth/react";
import { useEffect, useRef, useState } from "react";
import LoadingIconSrc from "../../../../public/loading-icon.png";
import Image from "next/image";

import { ChevronDownIcon, DoubleArrowLeftIcon } from "@radix-ui/react-icons";
import { redirect } from "next/dist/server/api-utils";

function useOutsideAlerter(ref, setFunction) {
  useEffect(() => {
    // Close Chevron dropdown when clicked outside
    function handleClickOutside(event) {
      if (ref.current && !ref.current.contains(event.target)) {
        setFunction(false);
      }
    }

    document.addEventListener("mousedown", handleClickOutside);

    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, [ref]);
}

export default function CollapsibleSidebar() {
  const { data: session, status } = useSession();
  const [authStatus, setAuthStatus] = useState("Untitled Workspace");
  const [chevronExpanded, setChevronExpanded] = useState(false);
  const dropdownRef = useRef(null);
  useEffect(() => {
    if (status === "authenticated") {
      const userName = `${session.user.name.split(" ")[0]}'s Workspace`;
      setAuthStatus(userName);
    } else if (status === "loading") {
      setAuthStatus("Loading...");
    }
  }, [status]);
  useOutsideAlerter(dropdownRef, setChevronExpanded);

  return (
    <div className="h-full border-r-[1px] bg-secondary">
      <div className="flex items-center justify-between gap-10 px-4 pt-4">
        <div className="flex items-center justify-center w-full gap-1">
          <Image
            src={
              status === "authenticated" ? session.user.image : LoadingIconSrc
            }
            alt="user-icon"
            width={36}
            height={36}
            className="mr-2 border-2 border-gray-400 rounded-md"
          />
          <p className="font-bold tracking-wide text-gray-600">{authStatus}</p>
          <button
            onClick={() => {
              setChevronExpanded(true);
            }}
          >
            <ChevronDownIcon />
          </button>
        </div>
        <DoubleArrowLeftIcon height={18} width={18} />
      </div>
      {chevronExpanded && (
        <div
          className="absolute w-[400px] mx-3 my-2 text-gray-600 rounded-md bg-background shadow-sm shadow-black"
          ref={dropdownRef}
        >
          <div className="pb-4 my-1 border-b-2 border-gray-300">
            <div className="px-4 py-1 text-sm font-semibold text-gray-500">
              {status === "authenticated" ? session.user.email : status}
            </div>
            <div className="flex gap-1 px-4 py-1">
              <Image
                src={
                  status === "authenticated"
                    ? session.user.image
                    : LoadingIconSrc
                }
                alt="user-icon"
                width={48}
                height={48}
                className="mr-2 border-2 border-gray-400 rounded-md"
              />
              <div className="flex flex-col">
                <p className="text-lg tracking-tight">{authStatus}</p>
                <div className="flex gap-3 text-xs">
                  <p className="text-gray-500">Free plan</p>
                  <p className="font-semibold text-gray-500">1 Members</p>
                </div>
              </div>
            </div>
          </div>
          <div className="grid gap-1 px-4 pb-4 mt-4">
            <span>Switch Workspace</span>
            <span>Account Settings</span>
            <span>
              <button onClick={() => redirect("/api/auth/signout")}>
                Log Out
              </button>
            </span>
          </div>
        </div>
      )}
    </div>
  );
}
