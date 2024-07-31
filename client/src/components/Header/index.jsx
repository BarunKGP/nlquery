import React from "react";
import { Button } from "../ui/button";
import Link from "next/link";

function Header() {
  return (
    <div>
      <div className="h-20 bg-headerbg w-full">
        <div className="grid grid-cols-12 gap-4 p-4 justify-items-center">
          <div className="text-themetext text-2xl col-span-2">nlQuery</div>
          <div className="text-themetext flex gap-4 col-span-7">
            <Link href={"/features"} className=" hover:underline px-4 py-2 ">
              Features
            </Link>
            <Link href={"/pricing"} className=" hover:underline px-4 py-2 ">
              Pricing
            </Link>
            <Link href={"/blog"} className=" hover:underline px-4 py-2 ">
              Blog
            </Link>
          </div>
          <div className="px-3 flex gap-3 tracking-tighter col-span-3">
            <Button className="bg-themetext hover:bg-purple-900">
              Request Demo
            </Button>
            <Button
              variant="secondary"
              className="bg-gray-800 text-gray-100 hover:text-gray-900"
            >
              Get Access
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Header;
