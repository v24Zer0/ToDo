import React, { useState } from "react";
import { Alert, Button, Text, TextInput, View } from "react-native";

type Props = {
    closeModal(): void;
}

const ItemCreateModal: React.FC<Props> = ({ closeModal }) => {
    const [task, setTask] = useState<string>("");
    const [priority, setPriority] = useState<string>("");

    return (
        <View>
            <Text>
                Create item
            </Text>
            <TextInput placeholder="Task" onChangeText={setTask} value={task} />
            <TextInput placeholder="Priority" onChangeText={setPriority} value={priority} />
            <Button title="Create" onPress={() => Alert.alert("Created item")} />
            <Button title="Cancel" onPress={closeModal} />
        </View>
    );
}

export default ItemCreateModal;