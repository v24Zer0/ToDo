import React from "react";
import { Button, Pressable, Text, View } from "react-native";
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
                {item.task}
            </Text>
            <Button title="Update item" onPress={ setUpdate } />
        </View>
    );
}

export default ItemModal;