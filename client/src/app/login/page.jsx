import GoogleButton from "./GoogleButton.server";
import GithubButton from "./GithubButton.server";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

export function Page() {
  return (
    <Card className="w-[500px] mx-auto mt-[200px]">
      <CardHeader>
        <CardTitle className="text-2xl">Sign in</CardTitle>
        <CardDescription>Log in with your provider</CardDescription>
      </CardHeader>
      <CardContent>
        <div className="flex flex-col items-center w-full gap-4">
          <GoogleButton />
          <GithubButton />
        </div>
      </CardContent>
      {/* <CardFooter className="flex justify-between">
        <Button variant="outline">Cancel</Button>
        <Button>Deploy</Button>
      </CardFooter> */}
    </Card>
  );
}

export default Page;
