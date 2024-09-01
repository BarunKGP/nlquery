export default function DashboardContents(props) {
  return (
    <div className="flex gap-8 py-1">
      <button className="text-sm tracking-wide rounded gradient-border-parent text-text">
        <div className="px-2 py-1 rounded-sm gradient-border-child">All</div>
      </button>
      <button className="text-sm tracking-wide rounded text-text gradient-border-parent">
        <div className="px-2 py-1 rounded-sm gradient-border-child">
          Recents
        </div>
      </button>
      <button className="text-sm tracking-wide text-text">Created by me</button>
      <button className="text-sm tracking-wide text-text">Folders</button>
      <button className="text-sm tracking-wide text-text">Templates</button>
    </div>
  );
}
