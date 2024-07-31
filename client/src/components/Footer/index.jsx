import Link from "next/link";

export default function Footer() {
  return (
    <div className="bg-headerbg text-themetext h-32 w-full">
      <div className="flex justify-between px-4 py-6">
        <div className="text-gray-500 hover:text-gray-400 font-semibold">
          <Link href="mailto:info@nlquery.com">info@nlquery.com</Link>
        </div>
        <div className="flex gap-12 px-2">
          <p className="text-gray-500 hover:text-gray-400">Privacy Policy</p>
          <p className="text-gray-500 hover:text-gray-400">Terms of Use</p>
          <p className="text-gray-500 hover:text-gray-400">
            <span> &copy; </span>
            2024 nlQuery
          </p>
        </div>
      </div>
    </div>
  );
}
