import React from "react";
import { Alert, Button, Text, View } from "react-native";
import Item from "../models/item";
import List from "../models/list";

interface Props {
    item: Item;
    list: List;
    setUpdate(): void;
}

const ItemModal: React.FC<Props> = ({ item, list, setUpdate }) => {
    return (
        <View>
            <Text>
                Update item
            </Text>
            <Text>
                {item.task}
            </Text>
            <Text>
                {item.priority}
            </Text>
            <Button title="Update item" onPress={ setUpdate } />
            <Button title="Delete item" onPress={() => Alert.alert("Item deleted")} />
        </View>
    );
}

export default ItemModal;