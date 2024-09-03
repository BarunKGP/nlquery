"use client";
import React from "react";
import { Button } from "../ui/button";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useSession } from "next-auth/react";
import Image from "next/image";

function Header() {
  const router = useRouter();
  const session = useSession();

  return (
    <div className="pt-2 pb-4 bg-headerbg">
      <div className="flex justify-center gap-10 px-2 py-1 mx-auto align-middle border-2 rounded-xl border-dgreen3 w-fit">
        <div
          className="flex justify-center gap-1 my-auto hover:opacity-100 opacity-70 hover:cursor-pointer"
          onClick={() =>
            session
              ? router.push(`/cards/${session.data.user.id}`)
              : router.push("/")
          }
        >
          <span className="italic text-textcolor">nl</span>
          <span className="text-themetext">Query</span>
          <span className="px-1 text-xs italic tracking-wider border-2 rounded-md border-textmuted h-fit text-textmuted">
            BETA
          </span>
        </div>
        <Link
          href={"/features"}
          className="px-4 py-2 hover:text-textcolor text-textmuted"
        >
          Features
        </Link>
        <Link
          href={"/pricing"}
          className="px-4 py-2 hover:text-textcolor text-textmuted"
        >
          Pricing
        </Link>
        <Link
          href={"/blog"}
          className="px-4 py-2 hover:text-textcolor text-textmuted"
        >
          Usecases
        </Link>
        {session && session.data && session.data.user ? (
          <div className="my-auto">
            <Image
              src={session.data.user.image}
              alt="user-profile-image"
              height={36}
              width={36}
              className="border-2 border-gray-400 rounded-full hover:border-themetext"
            />
          </div>
        ) : (
          <div>
            <Button className="border-2 bg-headerbd hover:text-textcolor hover:bg-dgreen1 text-textmuted border-dgreen3">
              <div>
                <Link href={"/login"}>
                  <span className="mr-2">Log In</span>
                  {/* <span className="px-2 py-1 border-2 rounded-md bg-dgreen3 border-dgreen1">
                  L
                </span> */}
                </Link>
              </div>
            </Button>
            <Button>Sign Up</Button>
          </div>
        )}
      </div>
    </div>
  );
}

export default Header;
