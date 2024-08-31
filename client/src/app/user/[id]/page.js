import CollapsibleSidebar from "./sidebar";

function Page() {
  return (
    <div className="flex w-full h-screen gap-4 bg-background">
      <CollapsibleSidebar />
      <div className="pt-5 text-3xl font-semibold text-center">
        This is the user page
      </div>
    </div>
  );
}

export default Page;
