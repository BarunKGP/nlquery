import FeatureList from "@/components/feature-list";

export default function TestSvg() {
  return (
    <div className="h-screen p-4">
      <FeatureList listKeys={["connector", "query", "dashboard", "share"]} />
    </div>
  );
}
