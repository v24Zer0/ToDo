import React from "react";
import { Alert, Button, Text, TextInput, View } from "react-native";

type Props = {
    setModalVisible(): void;
}

const ItemCreateModal: React.FC<Props> = ({ setModalVisible }) => {
    return (
        <View>
            <Text>
                Create item
            </Text>
            <TextInput />
            <TextInput />
            <Button title="Create" onPress={() => Alert.alert("Created item")} />
            <Button title="Cancel" onPress={setModalVisible} />
        </View>
    );
}

export default ItemCreateModal;