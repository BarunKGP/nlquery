"use client";
import githubSignInUrl from "../../../public/signin-assets/signin-assets-github/github-mark.svg?url";
import Image from "next/image";
import { signIn } from "next-auth/react";

export default function GoogleButton() {
  return (
    <button
      onClick={() => {
        signIn("github", { redirectTo: "/dashboard" });
      }}
    >
      <div className="flex gap-4">
        <Image
          src={githubSignInUrl}
          alt="github-sign-in"
          height={20}
          width={27}
        />
        <span>Sign in with Github</span>
      </div>
    </button>
  );
}
