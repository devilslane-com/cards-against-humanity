<?php

namespace CAH\Models;

use MongoDB\Laravel\Relations;

class Game extends BaseModel
{
    protected $casts = [
        'started_at' => 'datetime',
        'ended_at' => 'datetime',
    ];

    public function banned () : Relations\EmbedsMany
    {
        return $this->embedsMany (Player::class);
    }

    public function packs () : Relations\EmbedsMany
    {
        return $this->embedsMany (Pack::class);
    }

    public function players () : Relations\EmbedsMany
    {
        return $this->embedsMany (Player::class);
    }

    public function rounds () : Relations\HasMany
    {
        return $this->hasMany (Round::class);
    }

    public function winner () : Relations\HasOne 
    {
        return $this->hasOne (Player::class, 'winner');
    }
}
