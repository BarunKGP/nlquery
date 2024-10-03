import googleSignInUrl from "../../../public/signin-assets/signin-assets-google/Web (mobile + desktop)/svg/light/web_light_rd_na.svg?url";
import Image from "next/image";
import { signIn } from "@/auth";

export default async function GoogleButton() {
  return (
    <form
      action={async () => {
        "use server";
        await signIn("google", { redirectTo: "/dashboard" });
      }}
    >
      <button type="submit">
        <div className="flex items-center gap-2">
          <Image
            src={googleSignInUrl}
            alt="google-sign-in"
            height={24}
            width={24}
          />
          <span className="text-sm tracking-tighter text-text">
            Sign in with Google
          </span>
        </div>
      </button>
    </form>
  );
}
