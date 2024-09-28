"use client";

import { useSession } from "next-auth/react";
import { useCallback, useEffect, useRef, useState } from "react";
import NlQueryLogo from "../../../public/nlquery-logo.png";
import LoadingIconSrc from "../../../public/loading-icon.png";

import Image from "next/image";

import {
  ChevronDownIcon,
  DoubleArrowLeftIcon,
  DoubleArrowRightIcon,
} from "@radix-ui/react-icons";
import { redirect } from "next/navigation";
import Link from "next/link";

import ModeSwitcher from "@/components/mode-switcher";
import {
  SearchIcon,
  SquareKanbanIcon,
  ChartColumnStackedIcon,
  LayoutDashboardIcon,
  MessageSquareCodeIcon,
  Settings,
  SquareArrowUpRight,
  Star,
} from "lucide-react";
import { signOut } from "next-auth/react";

// Close Chevron dropdown when clicked outside
function useOutsideAlerter(ref, setFunction) {
  useEffect(() => {
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

const SidebarLinks = () => {
  return (
    <div className="grid grid-rows-3 mt-16 gap-14">
      <div className="flex flex-col gap-2 px-4">
        <div className="px-1 font-semibold text-text-muted hover:text-text hover:border-r-2 hover:border-text">
          <Link href={"/search"}>
            <div className="flex w-full gap-2">
              <span>
                <SearchIcon height={22} width={22} />
              </span>
              <span>Search</span>
            </div>
          </Link>
        </div>
        <div className="px-1 font-semibold text-text-muted hover:text-text hover:border-r-2 hover:border-text">
          <Link href={"/projects"}>
            <div className="flex gap-2">
              <span>
                <SquareKanbanIcon height={22} width={22} />
              </span>
              <span>Projects</span>
            </div>
          </Link>
        </div>
        <div className="px-1 font-semibold text-text-muted hover:text-text hover:border-r-2 hover:border-text">
          <Link href={"/dashboards"}>
            <div className="flex gap-2">
              <ChartColumnStackedIcon height={22} width={22} />
              <span>Dashboards</span>
            </div>
          </Link>
        </div>
        <div className="px-1 font-semibold text-text-muted hover:text-text hover:border-r-2 hover:border-text">
          <Link href={"/threads"}>
            <div className="flex gap-2">
              <MessageSquareCodeIcon width={22} height={22} />
              <span>Threads</span>
            </div>
          </Link>
        </div>
      </div>
      <div className="flex flex-col gap-2 px-4">
        <div className="px-1 text-text-muted hover:text-text hover:border-r-2 hover:border-text">
          <Link href={"/settings"} className="flex gap-2">
            <Settings color="#7383b5" strokeWidth={1} />
            <span className="text-[#7383b5]">Settings</span>
          </Link>
        </div>
        <div className="px-1 text-text-muted hover:text-text hover:border-r-2 hover:border-text">
          <Link href={"/shared"} className="flex gap-2">
            <SquareArrowUpRight color="#627f62" strokeWidth={1} />
            <span className="text-[#627f62]">Shared</span>
          </Link>
        </div>
        <div className="px-1 text-text-muted hover:text-text hover:border-r-2 hover:border-text">
          <Link href={"/favorites"} className="flex gap-2">
            <Star color="#ca9244" strokeWidth={1} />
            <span className="text-[#ca9244]">Favorites</span>
          </Link>
        </div>
        {/* <div className="px-1 font-semibold text-text-muted hover:text-text hover:border-r-2 hover:border-text">
          <Link href={"/threads"}>Threads</Link>
        </div> */}
      </div>
    </div>
  );
};

export default function CollapsibleSidebar() {
  const { data: session, status } = useSession();
  const [authStatus, setAuthStatus] = useState("Untitled Workspace");
  const [chevronExpanded, setChevronExpanded] = useState(false);
  const [isCollapsed, setIsCollapsed] = useState(false);
  const dropdownRef = useRef(null);

  const handleKeybind = useCallback((event) => {
    if (event.shiftKey && event.key.toLowerCase() === "s") {
      setIsCollapsed(() => (isCollapsed ? false : true));
    }
  });
  useEffect(() => {
    if (status === "authenticated") {
      const userName = `${session.user.name.split(" ")[0]}'s Workspace`;
      setAuthStatus(userName);
    } else if (status === "loading") {
      setAuthStatus("Loading...");
    }
  }, [status]);

  useEffect(() => {
    document.addEventListener("keydown", handleKeybind);
    return () => {
      document.removeEventListener("keydown", handleKeybind);
    };
  }, [handleKeybind]);

  useOutsideAlerter(dropdownRef, setChevronExpanded);

  return isCollapsed ? (
    <div className="absolute flex top-4 left-2">
      <div className="my-auto">
        <Image src={NlQueryLogo} width={60} height={40} />
      </div>
      <button
        className="p-1 my-6 rounded-full size-fit"
        onClick={() => setIsCollapsed(false)}
      >
        <DoubleArrowRightIcon
          width={16}
          height={16}
          className="dark:stroke-emerald-300 stroke-emerald-400"
        />
      </button>
    </div>
  ) : (
    <div className="h-full bg-secondary max-w-[240px]">
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
          <p className="font-bold text-text">{authStatus}</p>
          <button
            onClick={() => {
              setChevronExpanded(true);
            }}
          >
            <ChevronDownIcon />
          </button>
        </div>
        <button onClick={() => setIsCollapsed(true)}>
          <DoubleArrowLeftIcon height={16} width={16} />
        </button>
      </div>
      {chevronExpanded && (
        <div
          className="z-10 absolute w-[400px] mx-3 my-2 text-text rounded-md bg-background shadow-sm shadow-black"
          ref={dropdownRef}
        >
          <div className="pb-4 my-1 border-b-2 border-transparent">
            <div className="px-4 py-1 text-sm font-semibold text-text-muted">
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
                  <p className="text-text-muted">Free plan</p>
                  <p className="font-semibold text-text-muted">1 Members</p>
                </div>
              </div>
            </div>
          </div>
          <div className="grid gap-1 pb-4 mt-4">
            <span className="px-2 py-1 hover:bg-gray-200 dark:hover:bg-slate-800">
              Switch Workspace
            </span>
            <span className="px-2 py-1 hover:bg-gray-200 dark:hover:bg-slate-800">
              Account Settings
            </span>
            <span className="px-2 py-1 hover:bg-gray-200 dark:hover:bg-slate-800">
              Upgrade to Pro
            </span>
            <span className="px-2 py-1 hover:bg-gray-200 dark:hover:bg-slate-800">
              <button
                onClick={() => {
                  signOut({ callbackUrl: "/" });
                }}
              >
                Log Out
              </button>
            </span>
          </div>
        </div>
      )}
      <div className="m-2 mt-8 gradient-border-parent">
        <div className="flex items-center justify-center gap-1 px-2 gradient-border-child">
          <Image src={NlQueryLogo} width={60} height={40} />
          <span
            className="font-semibold tracking-tighter stippling text-text"
            data-text="nlQuery"
          >
            nlQuery
          </span>
        </div>
      </div>
      <SidebarLinks />
    </div>
  );
}
