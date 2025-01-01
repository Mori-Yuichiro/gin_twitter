"use client"

import Button from "@/components/button/Button";
import SignupModal from "@/components/modal/SignupModal";
import { useHomeHook } from "@/hooks/useHomeHook";

export default function Home() {
  const {
    openSignUpModal,
    setOpenSignUpModal
  } = useHomeHook();

  return (
    <>
      <main className="p-10 flex flex-col gap-y-12">
        <div>
          <svg className="mx-auto" xmlns="http://www.w3.org/2000/svg" width="20%" height="20%" viewBox="0 0 256 209"><path fill="#55acee" d="M256 25.45a105.04 105.04 0 0 1-30.166 8.27c10.845-6.5 19.172-16.793 23.093-29.057a105.183 105.183 0 0 1-33.351 12.745C205.995 7.201 192.346.822 177.239.822c-29.006 0-52.523 23.516-52.523 52.52c0 4.117.465 8.125 1.36 11.97c-43.65-2.191-82.35-23.1-108.255-54.876c-4.52 7.757-7.11 16.78-7.11 26.404c0 18.222 9.273 34.297 23.365 43.716a52.312 52.312 0 0 1-23.79-6.57c-.003.22-.003.44-.003.661c0 25.447 18.104 46.675 42.13 51.5a52.592 52.592 0 0 1-23.718.9c6.683 20.866 26.08 36.05 49.062 36.475c-17.975 14.086-40.622 22.483-65.228 22.483c-4.24 0-8.42-.249-12.529-.734c23.243 14.902 50.85 23.597 80.51 23.597c96.607 0 149.434-80.031 149.434-149.435c0-2.278-.05-4.543-.152-6.795A106.748 106.748 0 0 0 256 25.45" /></svg>
        </div>
        <div className="text-center">
          <h1 className="text-4xl">Gin + Next.js Twitterクローン</h1>
        </div>
        <div className="flex flex-col gap-y-3 items-center">
          <Button
            className="rounded-full bg-cyan-400 px-3 text-white hover:bg-cyan-600 w-1/4"
            onClick={setOpenSignUpModal}
          >アカウント作成</Button>
          <p>または</p>
          <Button className="rounded-full border border-black px-3 hover:bg-gray-400 hover:text-white w-1/4">ログイン</Button>
        </div>
      </main>
      {openSignUpModal && <SignupModal />}
    </>
  );
}
