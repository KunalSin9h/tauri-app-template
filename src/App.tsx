import Version from "./components/system/version";

export default function App() {
  return (
    <div className="p-4 w-full">
      <div className="flex text-4xl font-bold space-x-4 items-center justify-center">
        <span>Desktop App</span>
        <span>
          <Version />
        </span>
      </div>
    </div>
  );
}
