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
import { TableData } from "@/actions/get-table";

const formSchema = z.object({
  number_of_guests: z.coerce.number().min(1),
  table_number: z.coerce.number().min(1),
});

type TableFormValues = z.infer<typeof formSchema>;

interface TableFormProps {
  initialData: TableData | null;
}

export const TableForm: React.FC<TableFormProps> = ({ initialData }) => {
  const params = useParams();
  const router = useRouter();
  const storeId = params.storeId.toString();

  const [open, setOpen] = useState(false);
  const [loading, setLoading] = useState(false);

  const title = initialData ? "Edit table" : "Create table";
  const description = initialData ? "Edit a table." : "Add a new table";
  const toastMessage = initialData ? "Table updated." : "Table created.";
  const action = initialData ? "Save changes" : "Create";

  const defaultValues = initialData
    ? {
       ...initialData
      }
    : {
        number_of_guests: 0,
        table_number: 0,
        id: "",
      };

  const form = useForm<TableFormValues>({
    resolver: zodResolver(formSchema),
    defaultValues,
  });

  const onSubmit = async (data: TableFormValues) => {
    try {
      setLoading(true);
      if (initialData) {
        await axios.patch(
          `${process.env.NEXT_PUBLIC_API_URL}/${storeId}/table/${params.tableId}`,
          data
        );
      } else {
        await axios.post(
          `${process.env.NEXT_PUBLIC_API_URL}/${storeId}/table`,
          data
        );
      }
      router.refresh();
      router.push(`/${params.storeId}/table`);
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
      await axios.delete(`http://localhost:8080/table/${params.tableId}`);
      router.refresh();
      router.push(`/${params.storeId}/table`);
      toast.success("Table deleted.");
    } catch (error: any) {
      toast.error("Something went wrong.");
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
              name="table_number"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Table Number</FormLabel>
                  <FormControl>
                    <Input
                      type="number"
                      disabled={loading}
                      placeholder="Table Number"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="number_of_guests"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Number of Guests</FormLabel>
                  <FormControl>
                    <Input
                      type="number"
                      disabled={loading}
                      placeholder="Number of Guests (2-10)"
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
