import { LoginForm } from "@/components/form/authForm/login-form";

export const LoginPage = () => {
  return (
    <div className="h-screen w-full">
      <div className="flex h-full w-full">
        <div className="flex-1 flex justify-center items-center">
          <div className="bg-[#FE4E10] w-[90%] h-[90%] rounded-xl"></div>
        </div>
        <div className="flex-1 flex justify-center items-center h-full">
          <LoginForm className={"max-w-md w-full"} />
        </div>
      </div>
    </div>
  );
};
