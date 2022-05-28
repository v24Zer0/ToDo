import React, { useState } from "react";
import { Alert, Button, Text, View } from "react-native";
import { TextInput } from "react-native-gesture-handler";
import List from "../models/list";

interface Props {
    list: List;
    setUpdate(): void;
}

const ListModal: React.FC<Props> = ({ list, setUpdate }) => {
    const [name, setName] = useState<string>("");

    return (
        <View>
            <Text>
                Update list
            </Text>
            <TextInput placeholder={list.name} onChangeText={setName} value={name} />
            <Button title="Update list" onPress={ setUpdate } />
            <Button title="Delete list" onPress={() => Alert.alert("List deleted")} />
        </View>
    );
}

export default ListModal;