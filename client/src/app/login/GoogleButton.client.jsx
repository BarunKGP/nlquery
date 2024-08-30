"use client";
import googleSignInUrl from "../../../public/signin-assets/signin-assets/Web (mobile + desktop)/svg/light/web_light_rd_SI.svg?url";
import Image from "next/image";
import { signIn } from "next-auth/react";

export default function GoogleButton() {
  return (
    <button
      onClick={() =>
        signIn("google", {
          callbackUrl: "/cards/123",
        })
      }
    >
      <Image
        src={googleSignInUrl}
        alt="google-sign-in"
        height={72}
        width={180}
      />
    </button>
  );
}
