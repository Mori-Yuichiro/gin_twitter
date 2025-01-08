import Button from "../button/Button";

export default function RightSidebar() {
    return (
        <div className="min-h-screen w-1/4 pl-5 md:pl-6 xl:pl-8 pr-4 sm:pr-12 md:pr-20 xl:pr-28 flex flex-col gap-y-3 max-lg:hidden fixed right-0 border-l border-gray-300 box-border">
            <input
                type="text"
                placeholder="Search"
                className="border border-gray-200 rounded-full p-2"
            />
            <div className="flex flex-col gap-y-2 border border-black p-3 rounded-md">
                <h1 className="font-bold text-md xl:text-xl">Try Premium</h1>
                <p>Upgrade your experience with less ads, power tools, and more with Premium.</p>
                <Button className="border border-black rounded-full px-2 py-1">Start Free 14-day trial</Button>
            </div>
            <div className="border border-black rounded-md p-3 flex flex-col gap-y-3">
                <p>Golang</p>
                <p>Echo</p>
                <p>Gin</p>
                <p>gRPC</p>
                <p>Ruby on rails</p>
            </div>
        </div>
    );
}