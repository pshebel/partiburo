import { Party } from '../interfaces/party';
import { useAppForm } from '../hooks/form.tsx';
import { formOptions } from '@tanstack/react-form';

// Move your validation logic here so both Create and Edit share it
export const partyFormOptions = formOptions({
  defaultValues: {
    admin_email: '',
    title: '',
    description: '',
    address: '',
    date: '',
    time: '',
    reminders: [] as string[],
  },
  validators: {
    onChangeAsync: async ({ value }) => {
        const errors = {
            fields: {},
        } as {
            fields: Record<string, string>
        }
        if (!value.title) {
            errors.fields.title = 'Title is required'
        }
        if (!value.description) {
            errors.fields.description = 'Description is required'
        }
        if (!value.address) {
            errors.fields.address = 'Address is required'
        }
        if (!value.date) {
            errors.fields.date = 'Date is required'
        }

        if (!value.time) {
            errors.fields.time = 'Time is required'
        }

        if (value.title.length > 500) {
            errors.fields.name = 'Title is too long. Must be less than 500 characters.'
        }
        
        if (value.description.length > 5000) {
            errors.fields.description = 'Description is too long. Must be less than 5000 characters.'
        }

        return errors
        },
    }
});

interface PartyFormProps {
  initialValues?: Partial<Party>;
  onSubmit: (values: Party) => Promise<void>;
  buttonLabel: string;
}

export const PartyForm = ({ initialValues, onSubmit, buttonLabel }: PartyFormProps) => {
  const form = useAppForm({
    ...partyFormOptions,
    // Inject current values if they exist
    defaultValues: {
      ...partyFormOptions.defaultValues,
      ...initialValues,
    },
    onSubmit: async ({ value }) => {
      await onSubmit(value as Party);
    },
  });

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        e.stopPropagation();
        form.handleSubmit();
      }}
      className="space-y-6"
    >
      <form.AppField name="title" children={(field) => <field.TextField label="Title" />} />
      <form.AppField name="description" children={(field) => <field.TextArea label="Description" />} />
      <form.AppField name="address" children={(field) => <field.TextField label="Address" />} />
      <div className="grid grid-cols-2 gap-4">
        <form.AppField name="date" children={(field) => <field.DateField label="Date" />} />
        <form.AppField name="time" children={(field) => <field.TimeField label="Time" />} />
      </div>
      {/* Add your Reminders MultiSelect here as well */}
      
      <div className="pt-4">
        <form.AppForm>
          <form.SubscribeButton label={buttonLabel} />
        </form.AppForm>
      </div>
    </form>
  );
};