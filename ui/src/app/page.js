import { SearchBar } from "@/components/component/search-bar";
import FeatureList from "@/components/feature-list";
import HeroImage from "../../public/temp-hero-image.png";
import Image from "next/image";

export default function Home() {
  return (
    <div className="grid gap-12 min-h-screen">
      <div className="grid grid-cols-2 w-4/5 mx-auto gap-4 justify-items-center items-center">
        <div className="flex flex-col items-center justify-center gap-2 p-2">
          <h1 className="text-5xl tracking-tighter">
            <span className="italic font-thin mr-1">nl</span>Query
          </h1>
          <p className="text-2xl font-thin tracking-tight">
            Empowering business intelligence through natural language
          </p>
        </div>
        <div>
          <div className="border-2 border-gray-900">
            <Image
              src={HeroImage}
              width={400}
              height={400}
              alt="hero-image"
              objectFit="cover"
            />
          </div>
        </div>
      </div>
      <div className="text-themetext h-96 border-2 bg-headerbg border-headerbg shadow-md shadow-black p-4 text-center">
        <p className="text-3xl tracking-tight font-light">
          Talk to your data in 4 easy steps
        </p>
        <FeatureList listKeys={["connector", "query", "dashboard", "share"]} />
      </div>
    </div>
  );
}