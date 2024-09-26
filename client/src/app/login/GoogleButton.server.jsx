import googleSignInUrl from "../../../public/signin-assets/signin-assets-google/Web (mobile + desktop)/svg/light/web_light_rd_SI.svg?url";
import Image from "next/image";
import { signIn } from "@/auth";

export default async function GoogleButton() {
  return (
    <form
      action={async () => {
        "use server";
        await signIn("google", { redirectTo: "/user/123/" });
      }}
    >
      <button type="submit">
        <Image
          src={googleSignInUrl}
          alt="google-sign-in"
          height={64}
          width={240}
        />
      </button>
    </form>
  );
}
