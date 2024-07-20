export default function Footer() {
  return (
    <div className="bg-gradient-to-br from-headerbg to-gray-900 text-themetext h-32 w-full">
      <div className="flex justify-between px-4 py-6">
        <div className="text-gray-500 font-semibold">info@nlquery.com</div>
        <div className="flex gap-12 px-2">
          <p className="text-gray-500">Privacy Policy</p>
          <p className="text-gray-500">Terms of Use</p>
          <p className="text-gray-500">
            <span> &copy; </span>
            2024 nlQuery
          </p>
        </div>
      </div>
    </div>
  );
}
