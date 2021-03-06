import React, { useState } from "react";
import { Alert, Button, Text, View } from "react-native";
import { TextInput } from "react-native-gesture-handler";
import Item from "../models/item";
import List from "../models/list";

interface Props {
    item: Item;
    list: List;
    setUpdate(): void;
}

const ItemModal: React.FC<Props> = ({ item, list, setUpdate }) => {
    const [task, setTask] = useState<string>("");
    const [priority, setPriority] = useState<string>("");

    return (
        <View>
            <Text>
                Update item
            </Text>
            <TextInput placeholder={item.task} onChangeText={setTask} value={task} />
            <TextInput placeholder={item.priority.toString()} onChangeText={setPriority} value={priority} />
            <Button title="Update item" onPress={ setUpdate } />
            <Button title="Delete item" onPress={() => Alert.alert("Item deleted")} />
        </View>
    );
}

export default ItemModal;