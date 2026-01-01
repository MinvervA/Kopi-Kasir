import { z } from "zod";

export const loginSchema = z.object({
  email: z.string().trim().email("format email tidak valid!"),
  password: z.string().min(6, "Password tidak boleh kurang dari 6 karakter!"),
});

export const loginDefaultValues = {
  email: "",
  password: "",
};
