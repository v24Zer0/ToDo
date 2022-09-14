import React, { useState } from "react";
import { Alert, Button, Text, View } from "react-native";
import { TextInput } from "react-native-gesture-handler";

type Props = {
    closeModal(): void;
}

const ListCreateModal: React.FC<Props> = ({ closeModal }) => {
    const [name, setName] = useState<string>("");

    return (
        <View>
            <Text>
                Create list
            </Text>
            <TextInput placeholder="Name" onChangeText={setName} value={name} />
            <Button title="Create" onPress={() => Alert.alert("Created list")} />
            <Button title="Cancel" onPress={closeModal} />
        </View>
    );
}

export default ListCreateModal;