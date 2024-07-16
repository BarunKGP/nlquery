import { SearchBar } from "@/components/component/search-bar";
import Image from "next/image";

export default function Home() {
  return (
    <div className="flex min-h-screen flex-col justify-start gap-24 items-center p-24 border-4 ">
      <div className="flex flex-col items-center gap-2 p-2">
        <h1 className="text-5xl font-semibold">Leyndash</h1>
        <p className="text-2xl font-thin tracking-tighter">
          Empowering business intelligence through natural language
        </p>
      </div>
      <SearchBar />
    </div>
  );
}
