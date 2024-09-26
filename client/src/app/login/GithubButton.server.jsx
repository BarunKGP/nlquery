import githubSignInUrl from "../../../public/signin-assets/signin-assets-github/github-mark.svg?url";
import Image from "next/image";
import { signIn } from "@/auth";

export default async function GithubButton() {
  return (
    <form
      action={async () => {
        "use server";
        await signIn("github", { redirectTo: "/user/123/" });
      }}
    >
      <button type="submit">
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
    </form>
  );
}
