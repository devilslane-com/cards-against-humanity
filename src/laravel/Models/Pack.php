<?php

namespace CAH\Models;

use MongoDB\Laravel\Relations;

class Pack extends BaseModel
{
    public function cards () : Relations\HasMany
    {
        return $this->hasMany (Card::class);
    }
}
