import { SearchBar } from "@/components/component/search-bar";
import FeatureList from "@/components/feature-list";
import HeroImage from "../../public/temp-hero-image.png";
import Image from "next/image";

export default function Home() {
  return (
    <div className="grid min-h-screen gap-12 bg-[#EDF0EE]">
      <div className="grid items-center w-4/5 grid-cols-2 gap-4 mx-auto justify-items-center">
        <div className="flex flex-col items-center justify-center gap-2 p-2">
          <h1 className="text-5xl tracking-tighter">
            <span className="mr-1 italic font-thin">nl</span>Query
          </h1>
          <p className="text-2xl font-thin tracking-tight text-[#00735C]">
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
      <div className="p-4 text-center border-2 shadow-md text-themetext h-96 bg-headerbg border-headerbg shadow-black">
        <p className="text-3xl font-light tracking-tight">
          Talk to your data in 4 easy steps
        </p>
        <FeatureList listKeys={["connector", "query", "dashboard", "share"]} />
      </div>
    </div>
  );
}
