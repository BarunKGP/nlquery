import githubSignInUrl from "../../../public/signin-assets/signin-assets-github/github-mark.svg?url";
import Image from "next/image";
import { signIn, auth } from "@/auth";
import { getServerSession } from "next-auth";

export default async function GithubButton() {
  const session = await auth();
  return (
    <form
      action={async () => {
        "use server";
        await signIn("github", { redirectTo: "/dashboard" });

        //   await signIn("github", { redirect: false })
        //     .then(async (res) => {
        //       if (res.error) {
        //         console.error("Invalid credentials");
        //       } else {
        //         console.log(JSON.stringify(session, null, 2));
        //         if (session.user) {
        //           window.location.replace(
        //             `/user/${session.user.id ? session.user.id : 122313}`
        //           );
        //         }
        //       }
        //     })
        //     .catch((err) => {
        //       console.error(`Error occured: ${err}`);
        //     });
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
