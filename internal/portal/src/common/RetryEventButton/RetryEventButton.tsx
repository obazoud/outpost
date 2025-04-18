import React, { useCallback, useContext, useState, MouseEvent } from 'react';
import Button from "../Button/Button";
import { ReplayIcon } from "../Icons";
import { showToast } from "../Toast/Toast";
import { ApiContext } from '../../app';

interface RetryEventButtonProps {
    eventId: string;
    destinationId: string;
    disabled: boolean;
    loading: boolean;
    mutate: () => void;
}

const RetryEventButton: React.FC<RetryEventButtonProps> = ({
    eventId,
    destinationId,
    disabled,
    loading,
    mutate
}) => {
    const apiClient = useContext(ApiContext);
    const [ retryingId, setRetryingId ] = useState<null|string>(null);

    const retryEvent = useCallback(async (e: MouseEvent<HTMLButtonElement>) => {
        e.stopPropagation();
        setRetryingId(eventId);
        try {
          await apiClient.fetch(`destinations/${destinationId}/events/${eventId}/retry`,
            {
              method: "POST",
            }
          );
          showToast("success", "Retry successful.");
          mutate();
        } catch (error: any) {
          showToast("error", "Retry failed. " + `${error.message.charAt(0).toUpperCase() + error.message.slice(1)}`);
        }
          
        setRetryingId(null);
      }, [apiClient, destinationId, eventId, mutate]);

    return (
        <Button
          minimal
          onClick={(e) => retryEvent(e)}
          disabled={disabled || retryingId === eventId}
          loading={loading || retryingId === eventId}
        >
          <ReplayIcon />
        </Button>
    );
};

export default RetryEventButton;