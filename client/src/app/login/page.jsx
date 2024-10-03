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
    <Card className="max-w-[600px] min-w-[400px] mx-auto mt-[200px]">
      <CardHeader>
        <CardTitle className="text-2xl">Sign in</CardTitle>
        <CardDescription>Log in with your provider</CardDescription>
      </CardHeader>
      <CardContent>
        <div className="flex items-center justify-around w-full gap-4 mt-6">
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
