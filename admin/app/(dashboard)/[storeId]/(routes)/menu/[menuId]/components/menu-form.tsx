"use client";

import * as z from "zod";
import axios from "axios";
import { useState } from "react";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { toast } from "react-hot-toast";
import { Trash } from "lucide-react";
import { useParams, useRouter } from "next/navigation";

import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Separator } from "@/components/ui/separator";
import { Heading } from "@/components/ui/heading";
import { AlertModal } from "@/components/modals/alert-modal";
import { MenuData } from "@/actions/get-menu";

const formSchema = z.object({
  name: z.string(),
  category: z.string(),
});

type MenuFormValues = z.infer<typeof formSchema>;

interface MenuFormProps {
  initialData: MenuData | null;
}

export const MenuForm: React.FC<MenuFormProps> = ({ initialData }) => {
  const params = useParams();
  const router = useRouter();

  const [open, setOpen] = useState(false);
  const [loading, setLoading] = useState(false);

  const title = initialData ? "Edit menu" : "Create menu";
  const description = initialData ? "Edit a menu." : "Add a new menu";
  const toastMessage = initialData ? "Menu updated." : "Menu created.";
  const action = initialData ? "Save changes" : "Create";

  console.log(initialData);

const form = useForm<MenuFormValues>({
  resolver: zodResolver(formSchema),
  defaultValues: initialData
    ? { ...initialData }
    : {
        category: "",
        name: "",
      },
});

  const onSubmit = async (data: MenuFormValues) => {
    try {
      setLoading(true);
      if (initialData) {
        await axios.patch(
          `http://localhost:8080/menu/${params.menuId}`,
          data
        );
      } else {
        await axios.post(`http://localhost:8080/menu`, data);
      }
      router.refresh();
      router.push(`/${params.storeId}/menu`);
      toast.success(toastMessage);
    } catch (error: any) {
      toast.error("Something went wrong.");
    } finally {
      setLoading(false);
    }
  };

  const onDelete = async () => {
    try {
      setLoading(true);
      await axios.delete(`http://localhost:8080/menu/${params.menuId}`);
      router.refresh();
      router.push(`/${params.storeId}/menu`);
      toast.success("Menu deleted.");
    } catch (error: any) {
      toast.error(
        "Make sure you removed all categories using this menu first."
      );
    } finally {
      setLoading(false);
      setOpen(false);
    }
  };

  return (
    <>
      <AlertModal
        isOpen={open}
        onClose={() => setOpen(false)}
        onConfirm={onDelete}
        loading={loading}
      />
      <div className="flex items-center justify-between">
        <Heading title={title} description={description} />
        {initialData && (
          <Button
            disabled={loading}
            variant="destructive"
            size="sm"
            onClick={() => setOpen(true)}
          >
            <Trash className="h-4 w-4" />
          </Button>
        )}
      </div>
      <Separator />
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="space-y-8 w-full"
        >
          <div className="md:grid md:grid-cols-3 gap-8">
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Name</FormLabel>
                  <FormControl>
                    <Input
                      disabled={loading}
                      placeholder="Menu Name"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
          <div className="md:grid md:grid-cols-3 gap-8">
            <FormField
              control={form.control}
              name="category"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Category</FormLabel>
                  <FormControl>
                    <Input
                      disabled={loading}
                      placeholder="Menu Category"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
          <Button disabled={loading} className="ml-auto" type="submit">
            {action}
          </Button>
        </form>
      </Form>
    </>
  );
};
