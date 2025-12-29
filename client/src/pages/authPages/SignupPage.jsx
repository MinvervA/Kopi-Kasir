import { SignupForm } from "@/components/form/authForm/signupForm";

export const SignupPage = () => {
  return (
    <div className="h-screen w-full">
      <div className="flex h-full w-full">
        <div className="flex-1 flex justify-center items-center">
          <div className="bg-[#FE4E10] w-[90%] h-[90%] rounded-xl"></div>
        </div>
        <div className="flex-1 flex justify-center items-center h-full">
          <SignupForm className={"max-w-md w-full"} />
        </div>
      </div>
    </div>
  );
};
