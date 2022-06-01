import React from "react";
import { Alert, Button, Text, View } from "react-native";
import { TextInput } from "react-native-gesture-handler";

type Props = {
    setModalVisible(): void;
}

const ListCreateModal: React.FC<Props> = ({ setModalVisible }) => {
    return (
        <View>
            <Text>
                Create list
            </Text>
            <TextInput />
            <TextInput />
            <Button title="Create" onPress={() => Alert.alert("Created list")} />
            <Button title="Cancel" onPress={setModalVisible} />
        </View>
    );
}

export default ListCreateModal;